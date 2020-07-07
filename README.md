# cobraplugins
<img align="left" src="https://github.com/arrase/cobraplugins/blob/develop/assets/cobraplugins.jpg?raw=true">

### Add cobraplugins

    import (      
      "github.com/arrase/cobraplugins"
    )

### Add simple plugin to rootCmd

    func init() {
        rootCmd.AddCommand(cobraplugins.GetCmd(<YOUR_PLUGIN_PATH>, "MainCmd"))
    }
