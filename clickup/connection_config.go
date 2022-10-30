package clickup

import (
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type PluginConfig struct {
    Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
    "token": {Type: schema.TypeString},
}

func ConfigInstance() interface{} {
    return &PluginConfig{}
}

func GetConfig(connection *plugin.Connection) PluginConfig {
    if connection == nil || connection.Config == nil {
        return PluginConfig{}
    }

    config, _ := connection.Config.(PluginConfig)
    return config
}
