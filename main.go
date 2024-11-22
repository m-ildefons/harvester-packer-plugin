package main

import (
	"fmt"
	"os"

	harvester "github.com/harvester/packer-plugin-harvester/builder/harvester"
	version "github.com/harvester/packer-plugin-harvester/version"

	sdkplugin "github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	plugins := sdkplugin.NewSet()
	plugins.SetVersion(version.PluginVersion)
	plugins.RegisterBuilder(sdkplugin.DEFAULT_NAME, new(harvester.Builder))

	err := plugins.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}
