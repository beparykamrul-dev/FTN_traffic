package dns

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

type ProbeTarget struct {
	Name    string
	Address string
	Anycast bool
	Role    string
	Enabled bool
}

type Collector struct {
	Probe    Probe
	Interval time.Duration
}

func (c Collector) Run(ctx context.Context, portal *Portal, targets []ProbeTarget) {
	interval := c.Interval
	if interval <= 0 { interval = 30 * time.Second }
	c.collect(ctx, portal, targets)
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done(): return
		case <-t.C: c.collect(ctx, portal, targets)
		}
	}
}

func (c Collector) collect(ctx context.Context, portal *Portal, targets []ProbeTarget) {
	for _, target := range targets {
		if !target.Enabled || strings.TrimSpace(target.Address) == "" { continue }
		latency, err := c.Probe.Check(ctx, target.Address)
		status := NodeStatus{Name: target.Name, Address: target.Address, Anycast: target.Anycast, Healthy: err == nil, CheckedAt: time.Now()}
		if err == nil { status.LatencyMS = uint64(latency / time.Millisecond) }
		portal.UpsertNode(status)
	}
}

func ValidateTarget(t ProbeTarget) error {
	if strings.TrimSpace(t.Name) == "" { return fmt.Errorf("DNS target name is required") }
	if strings.TrimSpace(t.Address) == "" { return fmt.Errorf("DNS target address is required") }
	if _, _, err := net.SplitHostPort(t.Address); err != nil { return fmt.Errorf("DNS target %q must be host:port: %w", t.Name, err) }
	return nil
}
