package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
)

type VersionsPage struct {
	Files []string
}

func init() {
	flag.IntVar(&port, "port", 8080, "The port at which to serve http.")
	flag.StringVar(&host, "host", "0.0.0.0", "The host at which to serve http.")
}

func main() {
	flag.Parse()
	fmt.Printf("Initial configuration: [host: %s] [port: %d]\n", host, port)
	serveHTTP()
}

func serveHTTP() {
	serveStaticResources()
	// serveAjaxMethods(server)
	serveRequests()
	activateServer()
}

func serveStaticResources() {
	_, file, _, _ := runtime.Caller(0)
	here := filepath.Dir(file)
	static := filepath.Join(here, "/web")
	http.Handle("/", http.FileServer(http.Dir(static)))
}

func serveRequests() {
	http.HandleFunc("/gobrew/install", gobrewInstallHandler)
	http.HandleFunc("/versions", versionsHandler)
}

func gobrewInstallHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("scripts/gobrew_install.sh")
	if err != nil {
		fmt.Fprintf(w, "problem fetching install script %s", err)
		return
	}
	fmt.Fprintf(w, "%s", body)
}

func versionsHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir("web/files")
	fileNames := make([]string, len(files))
	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}
	versionsPage := &VersionsPage{Files: fileNames}
	versionsTemplate, _ := template.ParseFiles("templates/versions.html")
	versionsTemplate.Execute(w, versionsPage)
	// body, err := ioutil.ReadFile("templates/versions.html")
	// if err != nil {
	// 	fmt.Fprintf(w, "problem rendering versions %s", err)
	// 	return
	// }
	// fmt.Fprintf(w, "%s", body)
}

// func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
// 	t, _ := template.ParseFiles(tmpl + ".html")
// 	t.Execute(w, p)
// }

func activateServer() {
	fmt.Printf("Serving HTTP at: http://%s:%d\n", host, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

var (
	port int
	host string
)
