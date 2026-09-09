package dataplane

import (
	"errors"
	"strings"
)

var ErrBGPRuntimeCommandInvalid = errors.New("invalid BGP runtime command")

func ValidateBGPRuntimeCommand(command string) error {
	if strings.TrimSpace(command) == "" { return ErrBGPRuntimeCommandInvalid }
	if strings.ContainsAny(command, "\r\n\x00") { return ErrBGPRuntimeCommandInvalid }
	return nil
}
