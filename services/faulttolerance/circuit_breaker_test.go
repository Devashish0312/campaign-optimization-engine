package faulttolerance

import (
	"testing"
	"errors"
)

func TestCircuitBreakerExecute(t *testing.T) {
	cb := &CircuitBreaker{Failures: 0, Threshold: 3}

	action := func() error {
		return errors.New("Service failure")
	}

	for i := 0; i < 3; i++ {
		err := cb.Execute(action)
		if err == nil {
			t.Errorf("Expected error but got nil on attempt %d", i+1)
		}
	}

	// After 3 failures, the circuit breaker should be open
	err := cb.Execute(action)
	if err == nil || err.Error() != "circuit breaker open" {
		t.Errorf("Expected circuit breaker open error, got: %v", err)
	}
}
