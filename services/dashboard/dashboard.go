package dashboard

import (
	"net/http"
	"log"
	"campaign-optimization-engine/services/campaign"
)

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}

type DashboardService struct{}

func (d *DashboardService) Start() {
	http.HandleFunc("/dashboard", DashboardHandler)
	log.Println("Starting web dashboard on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	campaigns := campaign.GetActiveCampaigns()
	// Render a simple dashboard for campaigns
	for _, c := range campaigns {
		w.Write([]byte("Campaign ID: " + c.ID + "\n"))
	}
}
