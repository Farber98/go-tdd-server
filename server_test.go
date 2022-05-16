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

func assertResponse(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertResponseError(t testing.TB, got int, want int) {
	t.Helper()
	if got == 0 {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req

}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return req
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) getPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	str := &StubPlayerStore{
		map[string]int{
			"Yoni": 20,
			"Naju": 30,
		}, nil,
	}
	sv := &PlayerServer{str}

	t.Run("200GET /players/Yoni returns score", func(t *testing.T) {
		req := newGetScoreRequest("Yoni")
		resp := httptest.NewRecorder()

		sv.ServeHTTP(resp, req)
		assertResponse(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "20")

	})
	t.Run("200GET /players/Naju returns score", func(t *testing.T) {
		req := newGetScoreRequest("Naju")
		resp := httptest.NewRecorder()
		sv.ServeHTTP(resp, req)
		assertResponse(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "30")
	})
	t.Run("404GET /players/<missing-player>", func(t *testing.T) {
		req := newGetScoreRequest("Missing")
		resp := httptest.NewRecorder()
		sv.ServeHTTP(resp, req)
		assertResponseError(t, resp.Code, http.StatusNotFound)
	})

}

func TestPOSTPlayers(t *testing.T) {
	str := &StubPlayerStore{
		map[string]int{}, nil,
	}
	sv := &PlayerServer{str}
	t.Run("202POST records wins", func(t *testing.T) {
		req := newPostWinRequest("Jere")
		resp := httptest.NewRecorder()
		sv.ServeHTTP(resp, req)
		assertResponse(t, resp.Code, http.StatusAccepted)
		if len(str.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(str.winCalls), 1)
		}
	})
}
