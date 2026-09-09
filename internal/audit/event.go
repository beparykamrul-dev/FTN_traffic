package audit

import "time"

type Event struct {
	Actor string `json:"actor"`
	Action string `json:"action"`
	Resource string `json:"resource"`
	Approved bool `json:"approved"`
	At time.Time `json:"at"`
}
