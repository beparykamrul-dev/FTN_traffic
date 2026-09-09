package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
	"github.com/beparykamrul-dev/FTN_traffic/internal/httpapi"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func main() {
	addr := os.Getenv("FTN_TRAFFIC_ADDR")
	if addr == "" { addr = ":8080" }
	store := traffic.NewStore()
	dnsPortal := ftndns.NewPortal()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	targets := loadDNSTargets()
	if len(targets) > 0 {
		go ftndns.Collector{Interval: 30 * time.Second}.Run(ctx, dnsPortal, targets)
	} else {
		log.Printf("DNS health collector disabled: FTN_DNS_TARGETS is empty")
	}

	log.Printf("ftn-traffic listening on %s", addr)
	if err := http.ListenAndServe(addr, httpapi.Handler(store, dnsPortal)); err != nil { log.Fatal(err) }
}

func loadDNSTargets() []ftndns.ProbeTarget {
	raw := strings.TrimSpace(os.Getenv("FTN_DNS_TARGETS"))
	if raw == "" { return nil }
	var out []ftndns.ProbeTarget
	for _, item := range strings.Split(raw, ",") {
		parts := strings.SplitN(strings.TrimSpace(item), "=", 2)
		if len(parts) != 2 { log.Printf("ignoring invalid FTN_DNS_TARGETS item %q", item); continue }
		name, address := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		target := ftndns.ProbeTarget{Name: name, Address: address, Enabled: true}
		if err := ftndns.ValidateTarget(target); err != nil { log.Printf("ignoring DNS target: %v", err); continue }
		out = append(out, target)
	}
	return out
}
