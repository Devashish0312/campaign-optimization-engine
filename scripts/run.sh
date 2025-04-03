#!/bin/bash

# Set environment variables
export DATABASE_URL="http://localhost:5432"
export REDIS_URL="http://localhost:6379"
export MESSAGE_QUEUE_URL="http://localhost:3001"

# Start the application
go run cmd/main.go
