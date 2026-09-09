package main

import (
	"log"
	"net/http"
	"os"

	"github.com/beparykamrul-dev/FTN_traffic/internal/httpapi"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func main() {
	addr := os.Getenv("FTN_TRAFFIC_ADDR")
	if addr == "" { addr = ":8080" }
	store := traffic.NewStore()
	log.Printf("ftn-traffic listening on %s", addr)
	if err := http.ListenAndServe(addr, httpapi.Handler(store)); err != nil { log.Fatal(err) }
}
