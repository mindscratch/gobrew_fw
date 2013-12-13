package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
)

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
}

func gobrewInstallHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("scripts/gobrew_install.sh")
	if err != nil {
		fmt.Fprintf(w, "problem fetching install script %s", err)
		return
	}
	fmt.Fprintf(w, "%s", body)
}

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
