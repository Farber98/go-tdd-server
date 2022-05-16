package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Yoni's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Yoni", nil)
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		got := resp.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
