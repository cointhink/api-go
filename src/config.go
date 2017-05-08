package main

import (
	"io/ioutil"

	"github.com/hjson/hjson-go"
)

type Config struct {
	config map[string]interface{}
}

func (c Config) read(filename string) {
	var err error
	json, err := ioutil.ReadFile(filename)
	if err = hjson.Unmarshal(json, &c.config); err != nil {
		panic(err)
	}
}
