package main

import (
	"github.com/eirwin/briefly-meetings/api"
	"log"
	"net/http"
)

func main() {
	log.Print("starting meeting service...")
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8282", router))
}
