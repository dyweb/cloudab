// +nirvana:api=descriptors:"Descriptor"

package apis

import (
	descriptorsv1 "github.com/dyweb/cloudab/pkg/apis/v1/descriptors"
	"github.com/dyweb/cloudab/pkg/middlewares"

	def "github.com/caicloud/nirvana/definition"
)

// Descriptor returns a combined descriptor for APIs of all versions.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Description: "APIs",
		Path:        "/apis",
		Middlewares: middlewares.Middlewares(),
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Children: []def.Descriptor{
			descriptorsv1.Descriptor(),
		},
	}
}
