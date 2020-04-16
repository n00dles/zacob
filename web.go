package main

// barebones webserver running on port 8080
// for IOT Home atomation system
//
// This code is a very much a WIP as I learn GoLang
// and will probably not do very much until I know what I am doing...

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

// variable declarations
var tpl *template.Template
var title = "Zacob v0.1"
var devs = []device{}

const (
	STATIC_DIR = "/public/"
	PORT       = "8080"
)

type device struct {
	Id          string
	Name        string
	Typ         string
	Description string
}

type pageContent struct {
	Devices []device
	Title   string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func ShowDevices(w http.ResponseWriter, r *http.Request) {

}

func GetDevice(w http.ResponseWriter, r *http.Request) {

}

func CreateDevice(w http.ResponseWriter, r *http.Request) {

}

func UpdateDevice(w http.ResponseWriter, r *http.Request) {

}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {

}

func main() {

	// get a list of registered devices
	devs = getDevices("./data/", ".json")

	r := mux.NewRouter()

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/"))))

	r.HandleFunc("/", idx)
	r.HandleFunc("/api/devices", ShowDevices).Methods("GET")
	r.HandleFunc("/api/devices/{id}", GetDevice).Methods("GET")
	r.HandleFunc("/api/devices", CreateDevice).Methods("POST")
	r.HandleFunc("/api/devices/{id}", UpdateDevice).Methods("PUT")
	r.HandleFunc("/api/devices/{id}", DeleteDevice).Methods("DELETE")

	// websocket server
	http.Handle("/echo", websocket.Handler(echoHandler))

	log.Fatal(http.ListenAndServe(":"+PORT, r))

}

func idx(w http.ResponseWriter, req *http.Request) {
	var p pageContent
	p.Title = title
	p.Devices = devs
	err := tpl.ExecuteTemplate(w, "index.gohtml", p)
	if err != nil {
		fmt.Println(err)
	}
}
