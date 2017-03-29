package main

import (
	"log"
	"net/http"
	"github.com/dimfeld/httptreemux"
	api "github.com/caiobrentano/golden-roll-items/api"
)

func main() {
	addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	router.Handler(http.MethodPost, "/user/:psnId", &api.CreateDestinyUser{})
	// router.Handler(http.MethodGet, "/user/:psnId", &api.GetDestinyUser{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}