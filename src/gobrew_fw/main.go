package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
)

func init() {
	flag.IntVar(&port, "port", 8080, "The port at which to serve http.")
	flag.StringVar(&host, "host", "127.0.0.1", "The host at which to serve http.")
}

func main() {
	flag.Parse()
	fmt.Printf("Initial configuration: [host: %s] [port: %d]\n", host, port)
	serveHTTP()
}

func serveHTTP() {
	serveStaticResources()
	// serveAjaxMethods(server)
	activateServer()
}

func serveStaticResources() {
	_, file, _, _ := runtime.Caller(0)
	here := filepath.Dir(file)
	static := filepath.Join(here, "/web")
	http.Handle("/", http.FileServer(http.Dir(static)))
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
