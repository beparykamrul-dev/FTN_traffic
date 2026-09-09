package dataplane

import "errors"

var ErrFirewallUnauthorized = errors.New("firewall rule unauthorized")

func ValidateRule(r Rule) error {
	if !r.Authorized { return ErrFirewallUnauthorized }
	if r.ID == "" || r.Family == "" || r.Direction == "" || r.Action == "" { return errors.New("firewall rule fields are incomplete") }
	return nil
}

func ValidateRules(rules []Rule) error {
	for _, r := range rules { if err := ValidateRule(r); err != nil { return err } }
	return nil
}
