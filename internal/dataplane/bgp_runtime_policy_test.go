package dataplane

import "testing"

func TestValidateBGPRuntimeCommand(t *testing.T) {
	if ValidateBGPRuntimeCommand("show bgp summary") != nil { t.Fatal("valid command rejected") }
	if ValidateBGPRuntimeCommand(" ") != ErrBGPRuntimeCommandInvalid { t.Fatal("blank command accepted") }
	if ValidateBGPRuntimeCommand("show\nbgp") != ErrBGPRuntimeCommandInvalid { t.Fatal("newline command accepted") }
}
