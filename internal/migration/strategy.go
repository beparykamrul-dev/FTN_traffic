package migration

import "fmt"

var allowedStrategies = map[string]bool{"ecmp":true,"weighted":true,"canary":true,"failover":true,"blue_green":true}

func ValidateStrategy(strategy string) error {
	if !allowedStrategies[strategy] { return fmt.Errorf("unsupported migration strategy: %s", strategy) }
	return nil
}
