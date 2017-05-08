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

	log.Printf("cstruct %+v", c.config)
	json, err := json.Marshal(c.config)
	log.Printf("json %s", json)

	c.parser, err = gojq.NewStringQuery(string(json))
	log.Printf("c.parser set %+v", c.parser)
	return err
}

func (c *Config) queryString(query string) string {
	log.Printf("c.parser is %+v", c.parser)
	response, err := c.parser.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return response.(string)
}
