package latency

import (
	"errors"
	"math"
)

var ErrInvalidMeasurement = errors.New("invalid latency measurement")

func ValidateMeasurement(m Measurement) error {
	values := []float64{m.RTTP50MS,m.RTTP95MS,m.JitterMS,m.LossPercent,m.Availability}
	for _, v := range values { if math.IsNaN(v) || math.IsInf(v,0) { return ErrInvalidMeasurement } }
	if m.RTTP50MS < 0 || m.RTTP95MS < 0 || m.JitterMS < 0 || m.LossPercent < 0 || m.LossPercent > 100 || m.Availability < 0 || m.Availability > 1 { return ErrInvalidMeasurement }
	if m.RTTP50MS > m.RTTP95MS { return ErrInvalidMeasurement }
	return nil
}
