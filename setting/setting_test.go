package setting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = `
region: 
  - ap-northeast-1
  - ap-northeast-2

credentials:
  - credential:
    name: admin
    apikey: ABCDEFG
    secretkey: abcdefg
`

func TestLoadSettings(t *testing.T) {
	regions, credentials := LoadSettings([]byte(config))
	assert.Equal(t, []string{"ap-northeast-1", "ap-northeast-2"}, regions.R)
	assert.Equal(t, []Credential{{Name: "admin", Apikey: "ABCDEFG", Secretkey: "abcdefg"}}, credentials.C)
}
