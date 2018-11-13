package main

import (
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	tt := []struct {
		name    string
		request APIRequest
		status  int
	}{
		{
			name:    "Invalid Body",
			request: APIRequest{Body: "invalid"},
			status:  http.StatusBadRequest,
		},
		{
			name:    "Check name",
			request: APIRequest{Body: `{"name":"Footters"}`},
			status:  http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := Handler(tc.request)

			if res.StatusCode != tc.status {
				t.Errorf("Error status, expected %d, got %d", tc.status, res.StatusCode)
			}
		})
	}
}
