package dataplane

import "context"

type HealthGateBGPStub struct { RouterAdapter; Sessions []BGPSession; Err error }
func (b HealthGateBGPStub) SessionSummary(context.Context) ([]BGPSession,error) { return b.Sessions,b.Err }
