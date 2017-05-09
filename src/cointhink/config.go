package cointhink

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/elgs/gojq"
	"github.com/hjson/hjson-go"
)

type Config struct {
	config map[string]interface{}
	Parser *gojq.JQ
}

var C Config

func init() {
	C = Config{}
}

func (c *Config) Read(filename string) error {
	var err error
	humanjson, err := ioutil.ReadFile(filename)
	err = hjson.Unmarshal(humanjson, &c.config)
	if err != nil {
		return err
	}

	json, err := json.Marshal(c.config)

	c.Parser, err = gojq.NewStringQuery(string(json))
	return err
}

func (c *Config) QueryString(query string) string {
	response, err := c.Parser.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return response.(string)
}
