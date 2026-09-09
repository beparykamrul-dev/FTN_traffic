package dataplane

import "errors"

var ErrCommandNotAllowed = errors.New("local runtime command is not allowed")

var allowedRuntimeCommands = map[string]bool{
	"ip": true,
	"tc": true,
	"nft": true,
	"vtysh": true,
	"birdc": true,
	"gobgp": true,
}

func ValidateRuntimeCommand(name string) error {
	if !allowedRuntimeCommands[name] { return ErrCommandNotAllowed }
	return nil
}
