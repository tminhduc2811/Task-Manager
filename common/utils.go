package common

import (
	. "../models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
func ReadConfig(path string) (Config, error)  {
	var config Config
	source, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	return config, err
}