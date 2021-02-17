package descriptors

// import (
// 	"github.com/dyweb/cloudab/pkg/handlers/metrics"

// 	def "github.com/caicloud/nirvana/definition"
// )

// var (
// 	metricCtr metrics.Controller = metrics.New()
// )

// func init() {
// 	register([]def.Descriptor{{
// 		Path:        "/experiments/{experiment}/metrics",
// 		Definitions: []def.Definition{reportMetrics},
// 	},
// 	}...)
// }

// var reportMetrics = def.Definition{
// 	Method:      def.Get,
// 	Summary:     "Report metrics to cloudab server",
// 	Description: "Report custom metrics to the server",
// 	Function:    abcCtr.GetABConfig,
// 	Parameters:  []def.Parameter{},
// 	Results:     def.DataErrorResults("The metrics"),
// }
