package main

// barebones webserver running on port 8080 
// for IOT Home atomation system
// 
// This code is a very much a WIP as I learn GoLang 
// and will probably not do very much until I know what I am doing... 


import(
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type device struct{
	id string 
	name string
	typ string
	description string
}

type devices struct{
	dev []device
}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}
var devs = []device{}

func main() {


	devs = getDevices("./data/", ".json")
	//fmt.Println(devs)
	for i := range(devs) {
        d := devs[i]
        fmt.Println("Device:", d)
    }
	http.HandleFunc("/", idx)
	http.Handle("/assets/", http.StripPrefix("/assets",http.FileServer(http.Dir("public"))))
	//http.ListenAndServe(":8080",nil)
	
}

func idx(w http.ResponseWriter, req *http.Request){

	fmt.Println(devs) 
	type webData struct{
		Title(string)
		Name(string)
	}
	pd := webData{Title:"test", Name:"Mike"}
	err := tpl.ExecuteTemplate(w, "index.gohtml",pd)
	if err != nil{
		fmt.Println(err)
	}
}

