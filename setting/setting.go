package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Credential struct {
    Name string `yaml:"name"`
    Apikey string `yaml:"apikey"`
    Secretkey string `yaml:"secretkey"`
}

func LoadCredential() []Credential {
    var c []Credential
    buf, err := ioutil.ReadFile("setting.yaml")
    if err != nil {
	log.Fatal(err)
    }

    err = yaml.Unmarshal(buf, &c)
    if err != nil {
	log.Fatal(err)
    }

    return c
}
