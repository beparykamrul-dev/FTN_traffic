package dataplane

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidCommunity = errors.New("invalid BGP community")

func ValidateCommunity(value string) error {
	value = strings.TrimSpace(value)
	parts := strings.Split(value, ":")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return ErrInvalidCommunity
	}
	asn, err := strconv.ParseUint(parts[0], 10, 16)
	if err != nil || asn > 65535 {
		return ErrInvalidCommunity
	}
	v, err := strconv.ParseUint(parts[1], 10, 16)
	if err != nil || v > 65535 {
		return ErrInvalidCommunity
	}
	return nil
}

func ValidateCommunities(values []string) error {
	for _, value := range values {
		if err := ValidateCommunity(value); err != nil { return err }
	}
	return nil
}
