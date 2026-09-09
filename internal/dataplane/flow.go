package dataplane

import "context"

type FlowCollector interface {
	Name() string
	Health(context.Context) error
	Start(context.Context) error
	Stop(context.Context) error
}

type FlowRecord struct {
	SourcePrefix string `json:"source_prefix"`
	DestinationPrefix string `json:"destination_prefix"`
	SourceASN uint32 `json:"source_asn"`
	DestinationASN uint32 `json:"destination_asn"`
	Interface string `json:"interface"`
	Protocol uint8 `json:"protocol"`
	Bytes uint64 `json:"bytes"`
	Packets uint64 `json:"packets"`
}
