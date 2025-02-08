package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"prod"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServer `yaml:"http_server"`
	PostgresUrl string     `yaml:"postgres_url"`
}

func Mustload() *Config {
	var configpath string

	configpath = os.Getenv("CONFIG_PATH")

	if configpath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configpath = *flags
		if configpath == "" {
			log.Fatal("config path is not found")
		}
	}

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		log.Fatalf("config does not exist at path %s ", configpath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configpath, &cfg) // go get -u github.com/ilyakaznacheev/cleanenv
	if err != nil {
		log.Fatal("not able to read config file")
	}

	return &cfg
}
