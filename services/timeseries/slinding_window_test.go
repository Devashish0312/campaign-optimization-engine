package timeseries

import (
	"testing"
)

func TestSlidingWindowAverage(t *testing.T) {
	sw := SlidingWindow{Size: 3}

	sw.Add(1.0)
	sw.Add(2.0)
	sw.Add(3.0)

	avg := sw.Average()
	if avg != 2.0 {
		t.Errorf("Expected average to be 2.0, got %v", avg)
	}
}

func TestSlidingWindowEviction(t *testing.T) {
	sw := SlidingWindow{Size: 2}

	sw.Add(1.0)
	sw.Add(2.0)
	sw.Add(3.0)

	avg := sw.Average()
	if avg != 2.5 {
		t.Errorf("Expected average to be 2.5 after eviction, got %v", avg)
	}
}
