package conf

import "github.com/jasontconnell/conf"

type Config struct {
	ConnectionString string `json:"connectionString"`
	ProtobufLocation string `json:"protobufLocation"`
}

func LoadConfig(filename string) Config {
	var cfg Config
	conf.LoadConfig(filename, &cfg)
	return cfg
}
