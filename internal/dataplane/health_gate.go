package dataplane

import "context"

type HealthGate struct { Router RouterAdapter; BGP BGPAdapter; RequireRPKI bool; MaxPrefixes uint64 }

func (g HealthGate) Check(ctx context.Context) error {
	if g.Router == nil { return ErrBackendUnavailable }
	if err := g.Router.Health(ctx); err != nil { return err }
	if g.BGP == nil { return ErrBackendUnavailable }
	sessions, err := g.BGP.SessionSummary(ctx)
	if err != nil { return err }
	return ValidateBGPSessions(sessions, g.RequireRPKI)
}
