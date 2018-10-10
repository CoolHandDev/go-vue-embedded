package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./front-end/dist")
	mux := http.NewServeMux()
	fs := http.FileServer(box)
	mux.Handle("/", fs)

	srv := &http.Server{
		Addr:    "127.0.0.1:4022",
		Handler: mux,
	}

	fmt.Println("serving at", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handleDefault(w http.ResponseWriter, r *http.Request) {

}
