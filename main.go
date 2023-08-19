package main

import (
	"github.com/lyda/steampipe-plugin-oapigen/oapigen"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: oapigen.Plugin})
}
