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

var title = "Zacob v0.1"

type device struct{
	Id string 
	Name string
	Typ string
	Description string
}

type pageContent struct {
    Devices      []device
    Title	 string
}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}
var devs = []device{}

func main() {

	// get a list of registered devices
	devs = getDevices("./data/", ".json")
	
	http.HandleFunc("/", idx)
	http.Handle("/assets/", http.StripPrefix("/assets",http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080",nil)
	
}

func idx(w http.ResponseWriter, req *http.Request){
	var p pageContent
	p.Title = title
	p.Devices = devs
	err := tpl.ExecuteTemplate(w, "index.gohtml",p)
	if err != nil{
		fmt.Println(err)
	}
}

