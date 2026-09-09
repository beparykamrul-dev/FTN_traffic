package dataplane

import "context"

type BGPRouteGate struct { RequireRPKI bool; RequireBFD bool; RPKI RPKIState; BFD BFDChecker }

func (g BGPRouteGate) Check(ctx context.Context) error {
	if err := EnforceRPKI(g.RPKI, g.RequireRPKI); err != nil { return err }
	if g.RequireBFD { return RequireBFD(ctx, g.BFD) }
	return nil
}
