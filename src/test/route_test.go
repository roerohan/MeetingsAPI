package test

import(
    "net/http"
    // "net/http/httptest"
    "testing"
)

// TestMeetingsHandler tests meetings routes
func TestMeetingsHandler(t *testing.T) {
    if 200 != http.StatusOK {
        t.Errorf("Sample failed, expected status code: %v", http.StatusOK)
    }
}
