package lifecycle

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type C struct {
	Token string `yaml:"token"`
}

func Config() string {
	file, err := os.ReadFile("./config/env.yml")
	if err != nil {
		log.Fatalf("Cannot open env.yml file: %v", err)
	}
	config := &C{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("Cannot unmarshal env.yml file: %v", err)
	}

	return config.Token
}
