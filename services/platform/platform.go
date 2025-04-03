package platform

import (
	"math/rand"
)

// PlatformStats holds real-time data for a platform
type PlatformStats struct {
	CPC        float64
	Conversion float64
}

// PlatformService handles updates for platform stats
type PlatformService struct {
	PlatformStats map[string]PlatformStats
}

func NewPlatformService() *PlatformService {
	return &PlatformService{
		PlatformStats: make(map[string]PlatformStats),
	}
}

// UpdatePlatformStats simulates updating platform stats in real-time
func (ps *PlatformService) UpdatePlatformStats(platform string) {
	ps.PlatformStats[platform] = PlatformStats{
		CPC:        rand.Float64() * 2.0, // Random CPC between 0 and 2
		Conversion: rand.Float64() * 0.1, // Random CVR between 0 and 0.1
	}
}
