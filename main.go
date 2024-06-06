package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func sendFile(w http.ResponseWriter, r *http.Request) {
	file := "sample.txt"
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, file)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sendfile", sendFile)
	http.ListenAndServe("127.0.0.1:3333", r)
}
