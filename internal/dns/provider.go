package dns

import "context"

type ZoneRecord struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	TTL   uint32 `json:"ttl"`
	Value string `json:"value"`
}

type Provider interface {
	Name() string
	Health(context.Context) error
	ListRecords(context.Context, string) ([]ZoneRecord, error)
	UpsertRecord(context.Context, string, ZoneRecord) error
	DeleteRecord(context.Context, string, string, string) error
}
