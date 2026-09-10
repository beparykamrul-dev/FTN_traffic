package migration

import "errors"

var ErrInvalidMigrationPlan = errors.New("invalid migration plan")

type Plan struct {
	Request Request
	CanaryPercent uint8
	ObserveSeconds uint32
	RollbackOnHealthFailure bool
}

func (p Plan) Validate() error {
	if p.Request.ID == "" || p.Request.From == "" || p.Request.To == "" || p.Request.ApprovalID == "" || p.Request.From == p.Request.To { return ErrInvalidMigrationPlan }
	if err := ValidateStrategy(p.Request.Strategy); err != nil { return ErrInvalidMigrationPlan }
	if p.CanaryPercent == 0 || p.CanaryPercent > 100 || p.ObserveSeconds == 0 || !p.RollbackOnHealthFailure { return ErrInvalidMigrationPlan }
	return nil
}
