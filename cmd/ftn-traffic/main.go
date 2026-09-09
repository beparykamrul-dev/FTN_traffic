package main

import (
	"log"
	"net/http"
	"os"

	"github.com/beparykamrul-dev/FTN_traffic/internal/httpapi"
	ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func main() {
	addr := os.Getenv("FTN_TRAFFIC_ADDR")
	if addr == "" { addr = ":8080" }
	store := traffic.NewStore()
	dnsPortal := ftndns.NewPortal()
	log.Printf("ftn-traffic listening on %s", addr)
	if err := http.ListenAndServe(addr, httpapi.Handler(store, dnsPortal)); err != nil { log.Fatal(err) }
}
