package setting

import (
	"log"

	"gopkg.in/yaml.v2"
)

type Regions struct {
	R []string `yaml:"region"`
}

type Credentials struct {
	C []Credential `yaml:"credentials"`
}

type Credential struct {
	Name      string `yaml:"name"`
	Apikey    string `yaml:"apikey"`
	Secretkey string `yaml:"secretkey"`
}

type Metrics struct {
	M []string `yaml:"metrics"`
}

func LoadSettings(config []byte) (Regions, Credentials, Metrics) {
	var r Regions
	var c Credentials
	var m Metrics

	err := yaml.Unmarshal(config, &r)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(config, &c)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(config, &m)
	if err != nil {
		log.Fatal(err)
	}

	return r, c, m
}
