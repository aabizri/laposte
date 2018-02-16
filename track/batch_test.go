package track

import (
	"testing"
)

func TestClient_TrackBatch(t *testing.T) {
	if apiKey == "" || resourceDeliveredCode == "" {
		t.Fatalf("APIKey or resourceDeliveredCode absent, cannot test")
	}

	cl := New(apiKey)
	resp, err := cl.TrackBatch([]string{resourceDeliveredCode, "potato"})
	if err != nil {
		t.Error(err);
	}
	t.Logf("%#v", resp);
}

