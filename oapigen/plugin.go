// Package oapigen generates tables from OpenAPI documents.
package oapigen

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin defines the oapigen plugin.
func Plugin(_ context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-oapigen",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		SchemaMode:       plugin.SchemaModeDynamic,
		TableMapFunc:     PluginTables,
	}
	return p
}

type key string

const (
	// keyPath has been added to avoid key collisions
	keyPath key = "path"
)

// PluginTables defines all the tables from the OpenaAPI documents.
func PluginTables(ctx context.Context, d *plugin.TableMapData) (map[string]*plugin.Table, error) {
	// Initialize tables
	tables := map[string]*plugin.Table{}

	// Parse the passed OpenAPI document.
	docs, err := oapigenParse(ctx, d.Connection, d)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		docCtx := context.WithValue(ctx, keyPath, doc)
		base := filepath.Base(doc)

		tableData, err := tableCSV(docCtx, d.Connection)
		if err != nil {
			plugin.Logger(ctx).Error("oapigen.PluginTables", "create_table_error", err, "path", i)
			return nil, err
		}

		// Skip the table if the file is empty
		if tableData != nil {
			tables[base[0:len(base)-len(filepath.Ext(base))]] = tableData
		}
	}

	return tables, nil
}

func oapigenParse(ctx context.Context, connection *plugin.Connection, d *plugin.TableMapData) ([]string, error) {
	// Glob paths in config
	// Fail if no paths are specified
	Config := GetConfig(connection)
	if Config.Documents == nil {
		return nil, errors.New("documents must be configured")
	}

	// Gather file path matches for the glob
	var matches []string
	docs := Config.Documents
	for _, i := range docs {
		files, err := d.GetSourceFiles(i)
		if err != nil {
			plugin.Logger(ctx).Error("oapigen.oapigenParse", "failed to fetch absolute path", err, "path", i)
			return nil, err
		}
		matches = append(matches, files...)
	}

	// Sanitize the matches to ignore the directories
	var oapigenFilePaths []string
	for _, i := range matches {
		// Check if file or directory
		fileInfo, err := os.Stat(i)
		if err != nil {
			plugin.Logger(ctx).Error("oapigen.oapigenParse", "error getting file info", err, "path", i)
			return nil, err
		}

		// Ignore directories
		if fileInfo.IsDir() {
			continue
		}

		oapigenFilePaths = append(oapigenFilePaths, i)
	}

	return oapigenFilePaths, nil
}
