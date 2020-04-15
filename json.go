package main

// json functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/jsonq"
)

func getDevices(pathS string, ext string) []device {
	devs := []device{}
	var d device

	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filepath.Ext(path) == ext {
				//files = append(files, f.Name())
				//fmt.Println("opening file: data/" +f.Name())
				plan, _ := ioutil.ReadFile("data/" + f.Name())
				data := map[string]interface{}{}
				dec := json.NewDecoder(strings.NewReader(string(plan)))
				dec.Decode(&data)
				jq := jsonq.NewQuery(data)
				id, err := jq.String("id", "id")
				name, err := jq.String("id", "name")
				typ, err := jq.String("id", "typ")
				desc, err := jq.String("description")

				if err != nil {
					fmt.Println("Error:", err)
				}

				d.Id = id
				d.Name = name
				d.Typ = typ
				d.Description = desc

				devs = append(devs, d)
				//fmt.Println(d)
			}
		}
		return nil
	})
	return devs
}
