package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func loadConfig(path string) config {
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	var conf config

	if err := json.Unmarshal(dat, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
