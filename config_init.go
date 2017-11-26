package config_init

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func loadJsonObject(fileName string) {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	obj := make(map[string]interface{})

	json.Unmarshal(file, &obj)
	return obj
}

func initEnv(obj map[string]interface{}) {
	for k, v := range obj {
		os.Setenv(k, v)
	}
}

func configureEnv(fileName string) {
	initEnv(loadJsonObject(fileName))
}
