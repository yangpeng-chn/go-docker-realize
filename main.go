package main

import (
	"log"
	"net/http"

	"github.com/yangpeng-chn/go-docker-realize/app"
)

func main() {
	if err := http.ListenAndServe(":8888", app.NewMux()); err != nil {
		log.Println(err)
	}
}
