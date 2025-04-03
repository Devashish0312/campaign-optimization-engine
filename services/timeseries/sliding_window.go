package timeseries

import (
	"log"
)

type SlidingWindow struct {
	Size  int
	Items []float64
}

func NewSlidingWindow(size int) *SlidingWindow {
	return &SlidingWindow{Size: size}
}

func (sw *SlidingWindow) Add(item float64) {
	if len(sw.Items) >= sw.Size {
		sw.Items = sw.Items[1:]
	}
	sw.Items = append(sw.Items, item)
}

func (sw *SlidingWindow) Average() float64 {
	var sum float64
	for _, item := range sw.Items {
		sum += item
	}
	return sum / float64(len(sw.Items))
}
