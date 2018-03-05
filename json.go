package main

// json functions

import(
	"os"
	"path/filepath"
	"strings"
	"encoding/json"
	"github.com/jmoiron/jsonq"
	"io/ioutil"
	"fmt"
)

func getDevices(pathS string, ext string) []device {
	devs := []device{}

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
				id, err := jq.String("id","id")
				name, err := jq.String("id","name")
				typ, err := jq.String("id","typ")
				desc, err := jq.String("description")

				if err != nil {
					fmt.Println("Error:", err)
				}
				
				d := device{
					id: id,
					name: name,
					typ: typ,
					description: desc,
				}
				devs = append(devs, d)
				//fmt.Println(d)
			}
		}
		return nil
	})
	return devs
}
