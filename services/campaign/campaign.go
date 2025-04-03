package campaign

import (
	"log"
	"time"
	"container/heap"
)

type Campaign struct {
	ID          string
	Budget      float64
	TargetReach int
	TargetConv  int
	CPC         float64
	CVR         float64
	Deadline    time.Time
	PriorityScore float64
}

type CampaignQueue []*Campaign

func (cq CampaignQueue) Len() int { return len(cq) }

func (cq CampaignQueue) Less(i, j int) bool {
	return cq[i].PriorityScore > cq[j].PriorityScore
}

func (cq CampaignQueue) Swap(i, j int) {
	cq[i], cq[j] = cq[j], cq[i]
}

func (cq *CampaignQueue) Push(x interface{}) {
	*cq = append(*cq, x.(*Campaign))
}

func (cq *CampaignQueue) Pop() interface{} {
	old := *cq
	n := len(old)
	item := old[n-1]
	*cq = old[0 : n-1]
	return item
}

func NewCampaignService() *CampaignService {
	return &CampaignService{}
}

type CampaignService struct{}

func (c *CampaignService) StartProcessing() {
	log.Println("Campaign Service Started")
	// Create sample campaigns
	campaigns := []*Campaign{
		{ID: "1", Budget: 1000, TargetReach: 10000, TargetConv: 500, CPC: 0.5, CVR: 0.05, Deadline: time.Now().Add(1 * time.Hour)},
	}

	// Sort campaigns by priority
	campaignQueue := &CampaignQueue{}
	heap.Init(campaignQueue)

	for _, campaign := range campaigns {
		heap.Push(campaignQueue, campaign)
	}

	// Process top-priority campaign
	topCampaign := heap.Pop(campaignQueue).(*Campaign)
	log.Printf("Processing Campaign: %s with Budget: $%.2f", topCampaign.ID, topCampaign.Budget)
}
