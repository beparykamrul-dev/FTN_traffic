package migration

import "context"

type ProviderBackend interface { Name() string; Kind() string; Health(context.Context,string) error; Migrate(context.Context,string,string) error }
type Backend struct { Provider, KindName, CredentialRef string; Authorized bool }
func (b Backend) Name() string{return b.Provider}
func (b Backend) Kind() string{return b.KindName}
