package main

import (
	"fmt"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/config"
	"github.com/caicloud/nirvana/plugins/logger"
	"github.com/caicloud/nirvana/plugins/metrics"
	"github.com/caicloud/nirvana/plugins/reqlog"
	pversion "github.com/caicloud/nirvana/plugins/version"

	"github.com/dyweb/cloudab/pkg/apis"
	customconfig "github.com/dyweb/cloudab/pkg/config"
	"github.com/dyweb/cloudab/pkg/filters"
	"github.com/dyweb/cloudab/pkg/modifiers"
	"github.com/dyweb/cloudab/pkg/store"
	"github.com/dyweb/cloudab/pkg/version"
)

const (
	defaultServiceName = "cloudab"
)

func main() {
	// Print nirvana banner.
	fmt.Println(nirvana.Logo, nirvana.Banner)

	opt := config.NewDefaultOption()
	opt.Port = 9999

	// Create nirvana command.
	cmd := config.NewNamedNirvanaCommand(defaultServiceName, opt)

	// Create plugin options.
	metricsOption := metrics.NewDefaultOption() // Metrics plugin.
	reqlogOption := reqlog.NewDefaultOption()   // Request log plugin.
	versionOption := pversion.NewOption(        // Version plugin.
		defaultServiceName,
		version.Version,
		version.Commit,
		version.Package,
	)
	loggerOption := &logger.Option{
		Debug:          false,
		Level:          0,
		OverrideGlobal: true,
	}

	// Enable plugins.
	cmd.EnablePlugin(metricsOption, reqlogOption, versionOption, loggerOption)

	customConfig := &customconfig.Config{
		MongoURI: "mongodb://localhost:27017",
		DBName:   defaultServiceName,
	}
	cmd.AddOption("", customConfig)

	// Create server config.
	serverConfig := nirvana.NewConfig()

	// Configure APIs. These configurations may be changed by plugins.
	serverConfig.Configure(
		nirvana.Filter(filters.Filters()...),
		nirvana.Modifier(modifiers.Modifiers()...),
		nirvana.Descriptor(apis.Descriptor()),
	)

	// Set nirvana command hooks.
	cmd.SetHook(&config.NirvanaCommandHookFunc{
		PreServeFunc: func(config *nirvana.Config, server nirvana.Server) error {
			if err := store.Initialize(customConfig); err != nil {
				serverConfig.Logger().Fatal(err)
			}

			// Output project information.
			config.Logger().Infof("Package:%s Version:%s Commit:%s", version.Package, version.Version, version.Commit)
			return nil
		},
		// PostServeFunc: func(config *nirvana.Config, server nirvana.Server, err error) error {
		// 	// Output project information.
		// 	config.Logger().Infof("Disconnecting to mongo")

		// 	// TODO(gaocegege): Make it work.
		// 	return store.Disconnect()
		// },
	})

	// Start with server config.
	if err := cmd.ExecuteWithConfig(serverConfig); err != nil {
		serverConfig.Logger().Fatal(err)
	}
}
