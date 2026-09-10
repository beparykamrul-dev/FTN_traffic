package latency

import "testing"

func TestValidateMeasurement(t *testing.T) {
	if err:=ValidateMeasurement(Measurement{RTTP50MS:10,RTTP95MS:20,JitterMS:2,LossPercent:.2,Availability:.999}); err!=nil { t.Fatal(err) }
	if err:=ValidateMeasurement(Measurement{RTTP50MS:30,RTTP95MS:20,Availability:1}); err!=ErrInvalidMeasurement { t.Fatalf("got %v",err) }
	if err:=ValidateMeasurement(Measurement{Availability:1.1}); err!=ErrInvalidMeasurement { t.Fatalf("got %v",err) }
}
