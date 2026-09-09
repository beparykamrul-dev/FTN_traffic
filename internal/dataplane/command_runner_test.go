package dataplane

import (
	"context"
	"testing"
)

func TestLocalCommandRunnerRejectsUnknownCommand(t *testing.T) {
	_, err := (LocalCommandRunner{}).Run(context.Background(), "sh", "-c", "echo unsafe")
	if err == nil { t.Fatal("expected command rejection") }
}

func TestLocalCommandRunnerRejectsInvalidArguments(t *testing.T) {
	_, err := (LocalCommandRunner{}).Run(context.Background(), "ip", "route", "add\nunsafe")
	if err == nil { t.Fatal("expected argument rejection") }
}
