package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req

}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) getPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	str := &StubPlayerStore{
		map[string]int{
			"Yoni": 20,
			"Naju": 30,
		},
	}
	sv := &PlayerServer{str}

	t.Run("returns Yoni's score", func(t *testing.T) {
		req := newGetScoreRequest("Yoni")
		resp := httptest.NewRecorder()

		sv.ServeHTTP(resp, req)
		assertResponseBody(t, resp.Body.String(), "20")

	})
	t.Run("returns Naju's score", func(t *testing.T) {
		req := newGetScoreRequest("Naju")
		resp := httptest.NewRecorder()
		sv.ServeHTTP(resp, req)
		assertResponseBody(t, resp.Body.String(), "30")
	})
}
