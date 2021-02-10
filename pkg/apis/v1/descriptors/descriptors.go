package descriptors

import (
	def "github.com/caicloud/nirvana/definition"
)

// descriptors describe APIs of current version.
var descriptors []def.Descriptor

// register registers descriptors.
func register(ds ...def.Descriptor) {
	descriptors = append(descriptors, ds...)
}

// Descriptor returns a combined descriptor for current version.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Description: "v1 APIs",
		Path:        "/v1",
		Children:    descriptors,
	}
}
