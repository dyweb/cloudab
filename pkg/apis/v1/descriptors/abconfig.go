package descriptors

import (
	"github.com/dyweb/cloudab/pkg/apis/v1/converters"
	"github.com/dyweb/cloudab/pkg/handlers/abconfig"

	def "github.com/caicloud/nirvana/definition"
)

var (
	abcCtr abconfig.Controller = abconfig.New()
)

func init() {
	register([]def.Descriptor{{
		Path:        "/experiments/{experiment}/abconfig",
		Definitions: []def.Definition{getABConfig},
	},
	}...)
}

var getABConfig = def.Definition{
	Method:      def.Get,
	Summary:     "Get experiment's AB test configuration",
	Description: "Get the AB test configuration for the given experiment",
	Function:    abcCtr.GetABConfig,
	Parameters: []def.Parameter{
		{
			Source:      def.Path,
			Name:        "experiment",
			Description: "experiment id",
			Operators: []def.Operator{
				converters.ObjectID,
			},
		},
		{
			Source:      def.Query,
			Name:        "userID",
			Description: "user unique ID",
		},
	},
	Results: def.DataErrorResults("AB config"),
}
