package dataplane

import (
	"context"
	"testing"
)

func TestBFDSessionChecker(t *testing.T) {
	s := BFDSession{ID:"bfd-1", Local:"192.0.2.1", Remote:"192.0.2.2", MinRxMS:50, MinTxMS:50, Multiplier:3, Up:true, Authorized:true}
	if err := (BFDSessionChecker{Session:s}).CheckBFD(context.Background()); err != nil { t.Fatal(err) }
	s.Up = false
	if err := (BFDSessionChecker{Session:s}).CheckBFD(context.Background()); err != ErrBFDRequired { t.Fatalf("got %v", err) }
}
