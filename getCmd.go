package cobraplugins

import (
	"github.com/spf13/cobra"
	"plugin"
)

func GetCmd(pluginPath, cmdName string) *cobra.Command {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		panic(err)
	}
	b, err := p.Lookup(cmdName)
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("Init" + cmdName)
	if err == nil {
		f.(func())()
	}
	return *b.(**cobra.Command)
}
