package descriptors

import (
	"github.com/dyweb/cloudab/pkg/apis/v1/converters"
	"github.com/dyweb/cloudab/pkg/handlers/metrics"

	def "github.com/caicloud/nirvana/definition"
)

var (
	metricCtr metrics.Controller = metrics.New()
)

func init() {
	register([]def.Descriptor{{
		Path:        "/experiments/{experiment}/metrics",
		Definitions: []def.Definition{reportMetrics},
	},
	}...)
}

var reportMetrics = def.Definition{
	Method:      def.Create,
	Summary:     "Report metrics to cloudab server",
	Description: "Report custom metrics to the server",
	Function:    metricCtr.ReportMetrics,
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
		{
			Source:      def.Query,
			Name:        "metricName",
			Description: "metric name",
		},
		{
			Source:      def.Query,
			Name:        "value",
			Description: "metric value",
		},
	},
	Results: def.DataErrorResults("The metrics"),
}
