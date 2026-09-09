package flow

import "context"

type Record struct {
	SourceASN uint32 `json:"source_asn"`
	DestinationASN uint32 `json:"destination_asn"`
	Bytes uint64 `json:"bytes"`
	Packets uint64 `json:"packets"`
}

type Collector interface {
	Name() string
	Listen(ctx context.Context) error
}
