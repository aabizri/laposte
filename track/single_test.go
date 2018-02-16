package track

import (
	"testing"
	"os"
)

var (
	apiKey string = os.Getenv("LAPOSTE_TRACK_APIKEY")
	resourceDeliveredCode string = os.Getenv("LAPOSTE_TRACK_DELIVERED_RESOURCE_CODE")
)

func TestClient_Track(t *testing.T) {
	if apiKey == "" || resourceDeliveredCode == "" {
		t.Fatalf("APIKey or resourceDeliveredCode absent, cannot test")
	}

	cl := New(apiKey)
	resp, err := cl.Track(resourceDeliveredCode)
	if err != nil {
		t.Error(err);
	}
	t.Logf("%#v", resp);
}
