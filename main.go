package main

import (
	"fmt"
	"log"
	"net/http"
)

//go:generate esc -o frontend-assets.go -pkg main ./front-end/dist

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(FS(false))
	mux.Handle("/front-end/", fs)
	mux.Handle("/css/", http.StripPrefix("/css/", fs))
	//mux.Handle("/js/", http.StripPrefix("/js/", fs))
	mux.Handle("/img/", http.StripPrefix("/img/", fs))
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
