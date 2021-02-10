package config

type Config struct {
	MongoURI string `desc:"URL to the mongodb"`
	DBName   string `desc:"database name in mongodb"`
}
