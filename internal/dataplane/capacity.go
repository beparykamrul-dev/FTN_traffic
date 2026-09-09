package dataplane

import "errors"

var ErrInsufficientHeadroom = errors.New("capacity headroom below required threshold")

type Capacity struct {
	CapacityMbps float64 `json:"capacity_mbps"`
	UsedMbps float64 `json:"used_mbps"`
	ReservePercent float64 `json:"reserve_percent"`
}

func (c Capacity) AvailableMbps() float64 { return c.CapacityMbps - c.UsedMbps }

func (c Capacity) HasHeadroom() error {
	if c.CapacityMbps <= 0 || c.UsedMbps < 0 || c.UsedMbps > c.CapacityMbps { return ErrInsufficientHeadroom }
	reserve := c.CapacityMbps * c.ReservePercent / 100
	if c.AvailableMbps() < reserve { return ErrInsufficientHeadroom }
	return nil
}
