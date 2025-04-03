package bidding

import (
	"testing"
)

func TestBidCampaign(t *testing.T) {
	campaign := &Campaign{
		ID:        "1",
		Budget:    1000,
		Platform:  "Google",
		CPC:       0.5,
		CVR:       0.05,
	}

	decisionChan := make(chan string)
	go BidCampaign(campaign, decisionChan)

	decision := <-decisionChan
	if decision == "" {
		t.Error("Expected a bidding decision but got empty")
	}
}

func TestOptimizeBidding(t *testing.T) {
	campaigns := []*Campaign{
		{ID: "1", Budget: 1000, Platform: "Google", CPC: 0.5, CVR: 0.05},
		{ID: "2", Budget: 2000, Platform: "Meta", CPC: 0.4, CVR: 0.06},
	}

	OptimizeBidding(campaigns)
	// We expect bidding decisions to be printed (simulated in the test)
}
