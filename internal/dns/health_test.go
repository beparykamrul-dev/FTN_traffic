package dns

import (
	"context"
	"encoding/binary"
	"net"
	"testing"
	"time"
)

func TestProbeCheck(t *testing.T) {
	ln, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil { t.Fatal(err) }
	defer ln.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		buf := make([]byte, 512)
		n, addr, err := ln.ReadFrom(buf)
		if err != nil || n < 12 { return }
		response := make([]byte, 12)
		copy(response, buf[:12])
		binary.BigEndian.PutUint16(response[2:4], 0x8180)
		binary.BigEndian.PutUint16(response[6:8], 1)
		_, _ = ln.WriteTo(response, addr)
	}()

	p := Probe{Timeout: time.Second}
	latency, err := p.Check(context.Background(), ln.LocalAddr().String())
	if err != nil { t.Fatal(err) }
	if latency <= 0 { t.Fatal("expected positive probe latency") }
	<-done
}

func TestProbeRejectsUnavailableEndpoint(t *testing.T) {
	p := Probe{Timeout: 20 * time.Millisecond}
	_, err := p.Check(context.Background(), "127.0.0.1:1")
	if err == nil { t.Fatal("expected probe error") }
}
