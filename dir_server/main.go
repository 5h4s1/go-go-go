package main

import (
	"flag"
	"net/http"
	"log"
	"fmt"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		// fmt.Printf(r.URL.String())
		next.ServeHTTP(w, r)
	})
}

func main() {
	dir_server := flag.String("d", ".", "Directory to server")
	port := flag.String("p","8000", "Port to listen")

	flag.Parse()
	fmt.Println("Path: ", *dir_server)
	fmt.Println("Port: ", *port)
	fileServer := http.FileServer(http.Dir(*dir_server))

	http.Handle("/", logMiddleware(fileServer))
	if err := http.ListenAndServe(":" + *port, nil); err != nil {
        log.Fatal(err)
    }
}