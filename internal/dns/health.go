package dns

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// Probe performs a minimal DNS wire-protocol query. It is a health check, not a zone mutation.
type Probe struct {
	Timeout time.Duration
}

func (p Probe) Check(ctx context.Context, address string) (time.Duration, error) {
	timeout := p.Timeout
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	var idBytes [2]byte
	if _, err := rand.Read(idBytes[:]); err != nil { return 0, fmt.Errorf("generate dns probe id: %w", err) }
	id := binary.BigEndian.Uint16(idBytes[:])
	packet := make([]byte, 12)
	binary.BigEndian.PutUint16(packet[0:2], id)
	binary.BigEndian.PutUint16(packet[2:4], 0x0100)
	binary.BigEndian.PutUint16(packet[4:6], 1)
	packet = append(packet, 0, 0, 0, 1, 0, 0, 0, 0)

	start := time.Now()
	d := net.Dialer{Timeout: timeout}
	conn, err := d.DialContext(ctx, "udp", address)
	if err != nil { return 0, fmt.Errorf("dns probe %s: %w", address, err) }
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(timeout))
	if _, err := conn.Write(packet); err != nil { return 0, fmt.Errorf("dns probe write %s: %w", address, err) }
	response := make([]byte, 512)
	n, err := conn.Read(response)
	if err != nil { return 0, fmt.Errorf("dns probe read %s: %w", address, err) }
	if n < 12 || binary.BigEndian.Uint16(response[0:2]) != id {
		return 0, fmt.Errorf("dns probe %s: invalid response", address)
	}
	if response[2]&0x80 == 0 { return 0, fmt.Errorf("dns probe %s: response flag not set", address) }
	if binary.BigEndian.Uint16(response[6:8]) == 0 && response[3]&0x0f != 3 {
		return 0, fmt.Errorf("dns probe %s: no answer or NXDOMAIN", address)
	}
	return time.Since(start), nil
}
