package dns

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestProbeCheck(t *testing.T) {
	ln, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil { t.Fatal(err) }
	defer ln.Close()

	p := Probe{Timeout: time.Second}
	latency, err := p.Check(context.Background(), ln.LocalAddr().String())
	if err != nil { t.Fatal(err) }
	if latency <= 0 { t.Fatal("expected positive probe latency") }
}

func TestProbeRejectsUnavailableEndpoint(t *testing.T) {
	p := Probe{Timeout: 20 * time.Millisecond}
	_, err := p.Check(context.Background(), "127.0.0.1:1")
	if err == nil { t.Fatal("expected probe error") }
}
