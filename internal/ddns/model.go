package ddns

import "time"

type Record struct {
	ID string `json:"id"`
	Zone string `json:"zone"`
	Name string `json:"name"`
	Type string `json:"type"`
	Address string `json:"address"`
	TTL uint32 `json:"ttl"`
	Enabled bool `json:"enabled"`
	LastSeenAt time.Time `json:"last_seen_at"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type Update struct {
	Zone string `json:"zone"`
	Name string `json:"name"`
	Address string `json:"address"`
	TTL uint32 `json:"ttl"`
	Source string `json:"source"`
	Actor string `json:"actor"`
	RequestID string `json:"request_id"`
	ApprovalID string `json:"approval_id"`
}

type ProviderTarget struct {
	Provider string `json:"provider"`
	Zone string `json:"zone"`
	RecordName string `json:"record_name"`
	Enabled bool `json:"enabled"`
}
