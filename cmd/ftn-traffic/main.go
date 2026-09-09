package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/beparykamrul-dev/FTN_traffic/internal/httpapi"
	ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func main() {
	addr := os.Getenv("FTN_TRAFFIC_ADDR")
	if addr == "" { addr = ":8080" }
	store := traffic.NewStore()
	dnsPortal := ftndns.NewPortal()
	seedDNSPortal(dnsPortal)
	log.Printf("ftn-traffic listening on %s", addr)
	if err := http.ListenAndServe(addr, httpapi.Handler(store, dnsPortal)); err != nil { log.Fatal(err) }
}

func seedDNSPortal(p *ftndns.Portal) {
	now := time.Now()
	for _, e := range []ftndns.EngineStatus{
		{Name: "powerdns_enterprise", Role: "authoritative", Enabled: true, Healthy: true, CheckedAt: now},
		{Name: "technitium_dns", Role: "authoritative", Enabled: true, Healthy: true, CheckedAt: now},
		{Name: "coredns", Role: "authoritative", Enabled: true, Healthy: true, CheckedAt: now},
		{Name: "hickory_dns", Role: "authoritative", Enabled: true, Healthy: true, CheckedAt: now},
		{Name: "unbound", Role: "recursive", Enabled: true, Healthy: true, CheckedAt: now},
		{Name: "smartdns", Role: "recursive", Enabled: true, Healthy: true, CheckedAt: now},
		{Name: "dnsdist", Role: "edge", Enabled: true, Healthy: true, CheckedAt: now},
	} { p.UpsertEngine(e) }
}
