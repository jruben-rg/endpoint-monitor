package monitor

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPerformRequest(t *testing.T) {

	tests := []struct {
		timeout    int
		wantStatus int
	}{
		{0, http.StatusOK},
		{0, http.StatusTeapot},
		{0, http.StatusNotFound},
		{1, http.StatusRequestTimeout},
	}

	for _, test := range tests {

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if test.timeout > 0 {
				time.Sleep(time.Duration(test.timeout+1) * time.Second)
			}
			w.WriteHeader(test.wantStatus)
		}))

		defer server.Close()

		client := server.Client()
		client.Timeout = 1 * time.Second

		got := request(client, server.URL)
		if got != test.wantStatus {
			t.Errorf("Got status %d, Want; %d\n.", got, test.wantStatus)
		}
	}

}
