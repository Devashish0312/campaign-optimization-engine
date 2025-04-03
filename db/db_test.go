package db

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	// Assuming you have a function to test the DB connection.
	err := ConnectToDatabase() // This should connect to your DB and check the connection.
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
}
