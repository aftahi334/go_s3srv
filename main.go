package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	r := mux.NewRouter()

	r.HandleFunc("/{filename}", getFile).Methods(http.MethodGet)
	r.HandleFunc("/{filename}", sendFile).Methods(http.MethodPost)
	r.HandleFunc("/{filename}", deleteFile).Methods(http.MethodDelete)
	r.HandleFunc("/", getList).Methods(http.MethodGet)

	fmt.Println("Starting server at port 3333")
	http.ListenAndServe("127.0.0.1:3333", r)
}

func getFile(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	filename := vars["filename"]

	// Open the file for reading only
	file, err := os.Open(filename)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Copy the file content to the http.ResponseWriter
	_, err = io.Copy(writer, file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func sendFile(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	filename := vars["filename"]

	// Open a new file for writing only
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Copy the file from the request body to the new file
	_, err = io.Copy(file, request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func deleteFile(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	filename := vars["filename"]

	// Delete the file
	err := os.Remove(filename)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func getList(writer http.ResponseWriter, request *http.Request) {
	// Get the list of files in the current directory
	files, err := os.ReadDir(".")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of files to the http.ResponseWriter
	for _, file := range files {
		_, err := writer.Write([]byte(file.Name() + "\n"))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	writer.WriteHeader(http.StatusOK)
}
