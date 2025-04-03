package faulttolerance

import (
	"log"
	"time"
)

type CircuitBreaker struct {
	Failures    int
	Threshold   int
	LastFailure time.Time
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{
		Failures:  0,
		Threshold: 3,
	}
}

func (cb *CircuitBreaker) Start() {
	log.Println("Fault Tolerance (Circuit Breaker) Service Started")
	// Simulate service checks
	for i := 0; i < 5; i++ {
		err := cb.Execute(func() error {
			return nil
		})
		if err != nil {
			log.Println("Circuit is open")
		} else {
			log.Println("Circuit is closed")
		}
	}
}

func (cb *CircuitBreaker) Execute(action func() error) error {
	if cb.Failures >= cb.Threshold {
		return fmt.Errorf("circuit breaker open")
	}

	err := action()
	if err != nil {
		cb.Failures++
		cb.LastFailure = time.Now()
	} else {
		cb.Failures = 0
	}

	return err
}
