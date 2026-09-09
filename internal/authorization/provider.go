package authorization

import "time"

type Status string

const (
	Pending   Status = "pending"
	Connected Status = "connected"
	Expired   Status = "expired"
	Disabled  Status = "disabled"
)

type ProviderAuth struct {
	Provider      string    `json:"provider"`
	Status        Status    `json:"status"`
	Method        string    `json:"method"`
	AccountRef    string    `json:"account_ref,omitempty"`
	Scopes        []string  `json:"scopes,omitempty"`
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	LastCheckedAt time.Time `json:"last_checked_at,omitempty"`
	CredentialRef string    `json:"credential_ref,omitempty"`
}

func (a ProviderAuth) Active(now time.Time) bool {
	if a.Status != Connected { return false }
	return a.ExpiresAt.IsZero() || a.ExpiresAt.After(now)
}
