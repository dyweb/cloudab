package descriptors

import (
	"github.com/dyweb/cloudab/pkg/apis/v1/converters"
	"github.com/dyweb/cloudab/pkg/handlers/experiments"

	"github.com/caicloud/nirvana/definition"
	def "github.com/caicloud/nirvana/definition"
)

var (
	expCtr experiments.Controller = experiments.New()
)

func init() {
	register([]def.Descriptor{{
		Path: "/experiments",
		Definitions: []def.Definition{
			listExperiments,
			createExperiment,
		},
	}, {
		Path:        "/experiments/{experiment}",
		Definitions: []def.Definition{getExperiment},
	},
	}...)
}

var listExperiments = def.Definition{
	Method:      def.List,
	Summary:     "List experiments",
	Description: "Query a specified number of experiments and returns an array",
	Function:    expCtr.ListExperiments,
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
	Function:    expCtr.CreateExperiment,
	Parameters: []def.Parameter{
		definition.BodyParameterFor("JSON body to describe the new experiment"),
	},
	Results: def.DataErrorResults("An experiment"),
}

var getExperiment = def.Definition{
	Method:      def.Get,
	Summary:     "Get experiment",
	Description: "Get an experiment by id",
	Function:    expCtr.GetExperiment,
	Parameters: []def.Parameter{
		{
			Source:      def.Path,
			Name:        "experiment",
			Description: "experiment id",
			Operators: []def.Operator{
				converters.ObjectID,
			},
		},
	},
	Results: def.DataErrorResults("An experiment"),
}
