package traffic

import "time"

type Event struct {
	Provider string    `json:"provider"`
	Plane    string    `json:"plane"`
	POP      string    `json:"pop"`
	Bytes    uint64    `json:"bytes"`
	Packets  uint64    `json:"packets"`
	At       time.Time `json:"at"`
}
