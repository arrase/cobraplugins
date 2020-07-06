package cobraplugins

import (
	"github.com/spf13/cobra"
	"plugin"
)

func BuildCmd(pluginPath string) *cobra.Command {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		panic(err)
	}
	cfg, err := p.Lookup(PLUGIN_CFG_VAR_NAME)
	if err != nil {
		panic(err)
	}

	mainCmd, err := p.Lookup(cfg.(*cobraplugins.PluginCfg).Main)
	if err != nil {
		panic(err)
	}

	return *mainCmd.(**cobra.Command)
}
