package dashboard

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDashboardHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DashboardHandler)
	handler.ServeHTTP(rr, req)

	// Check if status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", rr.Code)
	}

	// Check if the dashboard content includes a campaign-related message
	expected := "Campaign Dashboard"
	if !contains(rr.Body.String(), expected) {
		t.Errorf("Expected to find %q in the response body, but got %v", expected, rr.Body.String())
	}
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}
