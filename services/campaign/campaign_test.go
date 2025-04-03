package campaign

import (
	"testing"
	"time"
)

func TestCampaignPriorityUpdate(t *testing.T) {
	campaign := &Campaign{
		ID:          "123",
		Budget:      1000,
		TargetReach: 100000,
		TargetConv:  1000,
		CPC:         0.5,
		CVR:         0.05,
		Deadline:    time.Now().Add(1 * time.Hour),
	}

	UpdatePriority(campaign)

	if campaign.PriorityScore <= 0 {
		t.Errorf("Expected priority score to be greater than 0, got %v", campaign.PriorityScore)
	}
}

func TestCampaignQueueSorting(t *testing.T) {
	campaign1 := &Campaign{PriorityScore: 10}
	campaign2 := &Campaign{PriorityScore: 20}
	campaign3 := &Campaign{PriorityScore: 15}

	queue := &CampaignQueue{campaign1, campaign2, campaign3}
	heap.Init(queue)

	if queue[0].PriorityScore != 20 {
		t.Errorf("Expected highest priority score to be 20, got %v", queue[0].PriorityScore)
	}
}
