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
				id, err := jq.String("id")
				name, err := jq.String("name")
				typ, err := jq.String("typ")
				desc, err := jq.String("description")
				ip, err := jq.String("ip")
				hash, err := jq.String("hash")
				active, err := jq.String("active")
				status, err := jq.String("status")
				if typ == "switch" && status == "0" {
					status = ""
				}

				if err != nil {
					fmt.Println("Error:", err)
				}

				d.Id = id
				d.Name = name
				d.Typ = typ
				d.Description = desc
				d.Ip = ip
				d.Hash = hash
				d.Active = active
				d.Status = status

				devs = append(devs, d)
				//fmt.Println(d)
			}
		}
		return nil
	})
	return devs
}
