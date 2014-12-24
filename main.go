package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	path := flag.String("path", ".", "path to static files")
	port := flag.String("port", "8080", "port number, default: 8080")
	flag.Parse()

	fmt.Printf("Super simple static server started.\nListening on 127.0.0.1:%s, serving path \"%s\".\n", *port, *path)
	fmt.Printf("---------------------------\n")

	panic(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", *port), http.FileServer(http.Dir(*path))))
}
