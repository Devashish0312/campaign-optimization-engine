package main

import (
	"log"
	"campaign-optimization-engine/config"
	"campaign-optimization-engine/services/bidding"
	"campaign-optimization-engine/services/campaign"
	"campaign-optimization-engine/services/predictive"
	"campaign-optimization-engine/services/messagequeue"
	"campaign-optimization-engine/services/timeseries"
	"campaign-optimization-engine/services/faulttolerance"
	"campaign-optimization-engine/services/dashboard"
	"campaign-optimization-engine/db"
	"campaign-optimization-engine/utils"
)

func main() {
	// Load Configurations
	config.LoadConfig()

	// Initialize services
	err := db.ConnectToDatabase()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Set up Redis connection
	redisClient := utils.InitializeRedis()

	// Start Campaign Processing
	campaignService := campaign.NewCampaignService()
	campaignService.StartProcessing()

	// Set up the message queue
	messageQueueService := messagequeue.NewMessageQueueService()
	messageQueueService.Start()

	// Start the web dashboard
	dashboardService := dashboard.NewDashboardService()
	dashboardService.Start()

	// Set up Fault Tolerance Mechanisms
	circuitBreaker := faulttolerance.NewCircuitBreaker()
	circuitBreaker.Start()

	// Run predictive analytics
	predictiveService := predictive.NewPredictiveService()
	predictiveService.Run()

	// Start bidding logic
	biddingService := bidding.NewBiddingService()
	biddingService.Start()

	log.Println("Campaign Optimization Engine Started")
}
