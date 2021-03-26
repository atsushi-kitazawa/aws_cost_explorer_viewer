package setting

import (
	"io/ioutil"
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
    Name string `yaml:"name"`
    Apikey string `yaml:"apikey"`
    Secretkey string `yaml:"secretkey"`
}

func LoadSettings() (Regions, Credentials) {
    var r Regions
    var c Credentials
    buf, err := ioutil.ReadFile("setting.yaml")
    if err != nil {
	log.Fatal(err)
    }

    err = yaml.Unmarshal(buf, &r)
    if err != nil {
	log.Fatal(err)
    }

    err = yaml.Unmarshal(buf, &c)
    if err != nil {
	log.Fatal(err)
    }

    return r, c
}
