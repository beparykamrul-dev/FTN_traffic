package ingest

import (
	"errors"
	"strings"

	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

var ErrInvalidEvent = errors.New("invalid traffic event")

func Normalize(e traffic.Event) (traffic.Event, error) {
	e.Provider = strings.TrimSpace(e.Provider)
	e.Plane = strings.TrimSpace(e.Plane)
	e.POP = strings.TrimSpace(e.POP)
	if e.Provider == "" || e.Plane == "" || e.POP == "" {
		return traffic.Event{}, ErrInvalidEvent
	}
	return e, nil
}
