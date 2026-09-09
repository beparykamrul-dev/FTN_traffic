package dataplane

import (
	"errors"
	"strings"
)

var ErrProductionPolicy = errors.New("invalid production dataplane policy")

type ProductionPolicy struct {
	RequireApproval bool
	RequireAuthorization bool
	FailClosed bool
	CapturePayload bool
}

func (p ProductionPolicy) Validate() error {
	if !p.RequireApproval || !p.RequireAuthorization || !p.FailClosed || p.CapturePayload {
		return ErrProductionPolicy
	}
	return nil
}

func ValidateResourceID(id string) error {
	if strings.TrimSpace(id) == "" || strings.ContainsAny(id, "\r\n\x00") { return ErrProductionPolicy }
	return nil
}
