package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Item struct {
	Name        string `json:"name"`
	CallbackURL string `json:"callback_url"`
	Time        string `json:"time"`
}

type Config struct {
	Tasks []Item `json:"tasks"`
}

var Conf *Config

func init() {
	fp, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(fp)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &Conf)

	if err != nil {
		log.Fatal(err)
	}
}
