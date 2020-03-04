package main

import (
	"log"

	"github.com/braddle/go-http-template/app"
	"github.com/gorilla/mux"
)

func main() {
	a := app.New(mux.NewRouter())
	log.Fatal(a.Run(":8080"))
}
