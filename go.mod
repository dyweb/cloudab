// go.mod example
//
// Refer to https://github.com/golang/go/wiki/Modules#gomod
// for detailed go.mod and go mod command documentation.
//
// module github.com/my/module/v3
//
// require (
//     github.com/some/dependency v1.2.3
//     github.com/another/dependency v0.1.0
//     github.com/additional/dependency/v4 v4.0.0
// )

module github.com/dyweb/cloudab

go 1.15

require (
	github.com/caicloud/nirvana v0.3.0-alpha.1.0.20210127083821-2f781c0f0aa7
	github.com/google/uuid v1.0.0 // indirect
	github.com/spaolacci/murmur3 v0.0.0-20180118202830-f09979ecbc72
	go.mongodb.org/mongo-driver v1.4.6
)
