package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/elgs/gojq"
	"github.com/hjson/hjson-go"
)

type Config struct {
	config map[string]interface{}
	parser *gojq.JQ
}

func (c *Config) read(filename string) error {
	var err error
	humanjson, err := ioutil.ReadFile(filename)
	err = hjson.Unmarshal(humanjson, &c.config)
	if err != nil {
		return err
	}

	json, err := json.Marshal(c.config)

	c.parser, err = gojq.NewStringQuery(string(json))
	return err
}

func (c *Config) queryString(query string) string {
	response, err := c.parser.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return response.(string)
}
