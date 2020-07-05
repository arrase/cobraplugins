# Simple Cobra plugin

## Build the plugin

    $ ./build.sh

## Create a cobra app from scratch

    cobra init test --pkg-name test.com/test

## Edit cmd/root.go

### Add cobraplugins

    import (
      ......
      
      "github.com/arrase/cobraplugins"
      
      ......
    )

### Add simple plugin to rootCmd

    func init() {
        ....
        rootCmd.AddCommand(cobraplugins.GetCmd(<YOUR_PLUGIN_PATH>, "MainCmd"))
    }
