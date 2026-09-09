package dataplane

import "context"

type CommandRunner interface { Run(context.Context, string, ...string) ([]byte, error) }

type LinuxBackend struct { Runner CommandRunner }

func (b LinuxBackend) Health(ctx context.Context) error {
	if b.Runner == nil { return ErrBackendUnavailable }
	_, err := b.Runner.Run(ctx, "ip", "-brief", "address")
	return err
}

func (b LinuxBackend) Routes(ctx context.Context) ([]byte, error) {
	if b.Runner == nil { return nil, ErrBackendUnavailable }
	return b.Runner.Run(ctx, "ip", "-j", "route", "show")
}

// Mutations intentionally remain outside this adapter until authorization and
// approval have been verified by the control plane.
