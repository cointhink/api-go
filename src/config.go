package main

import (
	"io/ioutil"

	"github.com/hjson/hjson-go"
)

type Config struct {
	config map[string]interface{}
}

func (c Config) read(filename string) error {
	var err error
	json, err := ioutil.ReadFile(filename)
	err = hjson.Unmarshal(json, &c.config)
	return err
}
