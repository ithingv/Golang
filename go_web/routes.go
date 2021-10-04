package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %s", r.URL.Path[1:])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

}
