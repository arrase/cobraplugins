package cobraplugins

const PLUGIN_CFG_VAR_NAME = "COBRA_PLUGIN_CFG"

type PluginCfg struct {
	Main   string
	Subs   []string
	Run    []string
}
