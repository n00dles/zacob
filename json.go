package main

// json functions

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getConfig() {
	jsondata, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(jsondata, &app)
}

func getDevices(pathS string, ext string) []device {
	devs = nil

	var d device

	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filepath.Ext(path) == ext {

				jsondata, _ := ioutil.ReadFile("data/" + f.Name())

				json.Unmarshal(jsondata, &d)

				devs = append(devs, d)

			}
		}
		return nil
	})
	return devs
}
