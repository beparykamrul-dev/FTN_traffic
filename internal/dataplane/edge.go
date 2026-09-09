package dataplane

import "context"

type EdgeBackend interface {
	Name() string
	Health(context.Context) error
	CacheStats(context.Context) (CacheStats, error)
	Purge(context.Context, string) error
}

type CacheStats struct {
	Requests uint64 `json:"requests"`
	Hits uint64 `json:"hits"`
	Misses uint64 `json:"misses"`
	OriginBytes uint64 `json:"origin_bytes"`
	EdgeBytes uint64 `json:"edge_bytes"`
}
