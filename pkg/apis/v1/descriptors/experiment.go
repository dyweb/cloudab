package descriptors

import (
	"github.com/dyweb/cloudab/pkg/handlers"
	"github.com/dyweb/cloudab/pkg/handlers/experiments"

	"github.com/caicloud/nirvana/definition"
	def "github.com/caicloud/nirvana/definition"
)

var (
	ctr experiments.Controller = experiments.New()
)

func init() {
	register([]def.Descriptor{{
		Path: "/experiments",
		Definitions: []def.Definition{
			listExperiments,
			createExperiment,
		},
	}, {
		Path:        "/experiments/{experiments}",
		Definitions: []def.Definition{getMessage},
	},
	}...)
}

var listExperiments = def.Definition{
	Method:      def.List,
	Summary:     "List experiments",
	Description: "Query a specified number of experiments and returns an array",
	Function:    ctr.ListExperiments,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "count",
			Default:     10,
			Description: "Number of experiments",
		},
	},
	Results: def.DataErrorResults("A list of experiments"),
}

var createExperiment = def.Definition{
	Method:      def.Create,
	Summary:     "Create experiment",
	Description: "Create a new experiment",
	Function:    ctr.CreateExperiment,
	Parameters: []def.Parameter{
		definition.BodyParameterFor("JSON body to describe the new experiment"),
	},
	Results: def.DataErrorResults("A list of experiments"),
}

var getMessage = def.Definition{
	Method:      def.Get,
	Summary:     "Get Message",
	Description: "Get a message by id",
	Function:    handlers.GetMessage,
	Parameters: []def.Parameter{
		def.PathParameterFor("message", "Message id"),
	},
	Results: def.DataErrorResults("A message"),
}
