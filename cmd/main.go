package main

import (
	"libraryassistant/api"
	"libraryassistant/service"
	"libraryassistant/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	path := "library.db"
	if v := os.Getenv("LIBRARY_DB"); v != "" {
		path = v
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	l := service.New(s)
	log.Println(http.ListenAndServe(":8080", api.New(l).Routes()))
}
