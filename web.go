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
var loggedIn = true
var devs = []device{}
var app config

type config struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Port      string `json:"port"`
	Staticdir string `json:"staticdir"`
	Debug     bool   `json:"debug"`
}

type device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Typ         string `json:"type"`
	Description string `json:"description"`
	IP          string `json:"ip"`
	Hash        string `json:"hash"`
	Active      string `json:"active"`
	Status      string `json:"status"`
}

type pageContent struct {
	Devices []device
	Title   string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	getConfig()
}

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// ShowDevices Show all devices
func ShowDevices(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Show Devices\n")
}

// GetDevice Get a single device
func GetDevice(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Get Device\n")
}

// CreateDevice Create a new device
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Create Device\n")
}

// UpdateDevice Update a device
func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Update Device\n")
}

// DeleteDevice Delete a device
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Delete Device\n")
}

// DoLogin Login page if not quthenticated
func DoLogin(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Login Page\n")
}

func init() {
	fmt.Println(app.Name + " " + app.Version)
	if loggedIn {
		fmt.Println("authenticated")
	}
}

func main() {

	// get a list of registered devices
	devs = getDevices("./data/", ".json")

	r := mux.NewRouter()

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("."+app.Staticdir))))

	r.HandleFunc("/", idx)
	r.HandleFunc("/login", DoLogin).Methods("GET", "POST")
	r.HandleFunc("/devices", ShowDevices).Methods("GET")
	r.HandleFunc("/devices/{id}", GetDevice).Methods("GET")
	r.HandleFunc("/devices", CreateDevice).Methods("POST")
	r.HandleFunc("/devices/{id}", UpdateDevice).Methods("PUT")
	r.HandleFunc("/devices/{id}", DeleteDevice).Methods("DELETE")

	// websocket server
	r.Handle("/echo/", websocket.Handler(echoHandler))

	log.Fatal(http.ListenAndServe(":"+app.Port, r))

}

func idx(w http.ResponseWriter, req *http.Request) {
	if app.Debug == true {
		tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	}
	//if loggedIn == true {
	var p pageContent
	p.Title = app.Name + " " + app.Version
	p.Devices = devs
	err := tpl.ExecuteTemplate(w, "index.gohtml", p)
	if err != nil {
		fmt.Println(err)
	}
	//} else {
	//	http.Redirect(w, req, "/login", http.StatusMovedPermanently)
	//}
}
