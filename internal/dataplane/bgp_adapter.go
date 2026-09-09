package dataplane

import "context"

type BGPAdapter interface {
	RouterAdapter
	SessionSummary(context.Context) ([]BGPSession, error)
}
