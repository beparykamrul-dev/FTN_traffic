package dns

import (
	"context"
	"fmt"
	"net"
	"time"
)

// Probe checks a DNS endpoint without issuing application-level record changes.
type Probe struct {
	Timeout time.Duration
}

func (p Probe) Check(ctx context.Context, address string) (time.Duration, error) {
	timeout := p.Timeout
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	start := time.Now()
	d := net.Dialer{Timeout: timeout}
	conn, err := d.DialContext(ctx, "udp", address)
	if err != nil {
		return 0, fmt.Errorf("dns probe %s: %w", address, err)
	}
	_ = conn.Close()
	return time.Since(start), nil
}
