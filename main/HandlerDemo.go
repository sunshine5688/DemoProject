package main

import (
	"log"
	"net/http"
)

type httpServer struct {
}

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func main() {
	var server httpServer
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}
