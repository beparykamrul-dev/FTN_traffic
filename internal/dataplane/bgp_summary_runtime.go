package dataplane

import "context"

func CollectBGPSessions(ctx context.Context, kind RouterKind, source BGPSummarySource, requireRPKI bool, maxPrefixes uint64) ([]BGPSession, error) {
	if source == nil {
		return nil, ErrBackendUnavailable
	}
	out, err := source.SummaryOutput(ctx)
	if err != nil {
		return nil, err
	}
	sessions, err := ParseBGPSummary(kind, out)
	if err != nil {
		return nil, err
	}
	if err := ValidateBGPSessionsWithLimit(sessions, requireRPKI, maxPrefixes); err != nil {
		return sessions, err
	}
	return sessions, nil
}
