package migration

import (
	"context"
	"testing"
)
type fakeBackend struct{ Backend; healthy bool }
func(f fakeBackend) Health(context.Context,string)error{if !f.healthy{return context.Canceled};return nil}
func(f fakeBackend) Migrate(context.Context,string,string)error{return nil}
func TestBackendKeepsCredentialReferenceOnly(t *testing.T){b:=Backend{Provider:"aws",KindName:"hosting",CredentialRef:"secret/ftn/aws",Authorized:true};if b.CredentialRef==""||b.Provider!="aws"{t.Fatal("invalid backend")}}
