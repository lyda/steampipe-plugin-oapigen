// Package oapigen generates tables from OpenAPI documents.
package oapigen

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func readOpenAPI(ctx context.Context, connection *plugin.Connection, path string) (*openapi3.T, error) {
	// TODO: Support other versions - this only supports OpenAPI v3.
	cfg := GetConfig(connection)
	if cfg.Version != 3 {
		err := fmt.Errorf("only OpenAPI version 3 is supported")
		plugin.Logger(ctx).Error("oapigen.readOpenAPI", "version_error", err)
		return nil, err
	}
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("oapigen.readOpenAPI", "loader_error", err, "path", path)
	}
	return doc, err
}

// tableOpenAPI creates the tables from an OpenAPI document.  Some useful docs:
// Spec: https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md
// Kin lib: https://pkg.go.dev/github.com/getkin/kin-openapi/openapi3
func tableOpenAPI(ctx context.Context, base string, connection *plugin.Connection) ([]*plugin.Table, error) {
	cfg := GetConfig(connection)
	tableDefs := []*plugin.Table{}
	tableNamePfx := base[0 : len(base)-len(filepath.Ext(base))]

	docFilename := ctx.Value(keyPath).(string)
	doc, err := readOpenAPI(ctx, connection, docFilename)
	if err != nil {
		plugin.Logger(ctx).Error("oapigen.tableOpenAPI", "read_openapi_error", err, "path", docFilename, "version", cfg.Version)
		return nil, fmt.Errorf("failed to load OpenAPI file %s: %v", docFilename, err)
	}

	for path, pathItem := range doc.Paths {
		if pathItem.Get == nil || pathItem.Get.Deprecated {
			continue
		}
		tableName := fmt.Sprintf("%s_%s", tableNamePfx, strings.ReplaceAll(path, "/", "_"))
		cols := []*plugin.Column{}
		colNames := []string{}
		keyCols := []*plugin.KeyColumn{}

		if pathItem.Get.Parameters != nil {
			for _, param := range pathItem.Get.Parameters {
				if param.Value == nil {
					continue
				}
				// TODO: make sure this can be matched.
				colName := "param_" + param.Value.Name
				colNames = append(colNames, colName)
				col := &plugin.Column{
					Name:        colName,
					Transform:   transform.FromField(helpers.EscapePropertyName(param.Value.Name)),
					Description: fmt.Sprintf("Param %s", param.Value.Name),
				}
				if param.Value.Schema == nil || param.Value.Schema.Value == nil {
					col.Type = proto.ColumnType_STRING
				} else {
					switch param.Value.Schema.Value.Type {
					case "integer":
						col.Type = proto.ColumnType_INT
					// TODO: More types.
					default:
						col.Type = proto.ColumnType_STRING
					}
				}
				if param.Value.Required {
					keyCols = append(keyCols, &plugin.KeyColumn{
						Name:    colName,
						Require: plugin.Required,
					})
				}
				cols = append(cols, col)
			}
			// TODO: Loop through respose and add those columns.
		}
		tableDefs = append(tableDefs, &plugin.Table{
			Name:        tableName,
			Description: fmt.Sprintf("%s endpoint from %s", path, docFilename),
			List: &plugin.ListConfig{
				KeyColumns: keyCols,
				Hydrate:    listFromEndpoint("TODO"),
			},
			Columns: cols,
		})
	}

	return tableDefs, nil

}

func listFromEndpoint(_ string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		return nil, fmt.Errorf("TODO")
	}
}
