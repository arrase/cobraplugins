package cobra-plugins

import (
	"plugin"
  "github.com/spf13/cobra"
)

func Load(rootCdm *cobra.Command, pluginDir string) {
	p, err := plugin.Open(pluginDir)
	if err != nil {
		panic(err)
	}

	b, err := p.Lookup("PluginCmd")
	if err != nil {
		panic(err)
	}

  rootCdm.AddCommand(*b.(**cobra.Command))
}
