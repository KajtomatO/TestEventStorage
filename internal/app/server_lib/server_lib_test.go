package server_lib

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var errTestError = errors.New("test error")

type StubDataStore struct {
	data       map[string]string
	storeCalls []string
}

func (s *StubDataStore) GetDataRecord(id string) (string, error) {
	data := s.data[id]

	if len(data) == 0 {
		return "", errTestError
	}

	return data, nil
}

func (s *StubDataStore) SetDataRecord(id string, data string) {
	s.storeCalls = append(s.storeCalls, id+data)
}

func TestGETData(t *testing.T) {
	store := StubDataStore{
		map[string]string{
			"KlientA": "ABccdF",
			"KlientB": "Trolololo",
		},
		nil,
	}
	server := &DataServer{&store}

	tests := []struct {
		name               string
		id                 string
		expectedHTTPStatus int
		expecteddata       string
	}{
		{
			name:               "Returns KlientA data",
			id:                 "KlientA",
			expectedHTTPStatus: http.StatusOK,
			expecteddata:       "ABccdF",
		},
		{
			name:               "Returns KlientB data",
			id:                 "KlientB",
			expectedHTTPStatus: http.StatusOK,
			expecteddata:       "Trolololo",
		},
		{
			name:               "Returns 404 on missing client",
			id:                 "Mike",
			expectedHTTPStatus: http.StatusNotFound,
			expecteddata:       "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := getDataRequest(tt.id)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStatus)
			assertResponseBody(t, response.Body.String(), tt.expecteddata)
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := StubDataStore{
		map[string]string{},
		nil,
	}
	server := &DataServer{&store}

	t.Run("it saves data with POST", func(t *testing.T) {
		user := "Arnold"
		data := "Hello world!"

		request := newPostSaveRequest(user, data)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.storeCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.storeCalls), 1)
		}

		if store.storeCalls[0] != user+data {
			t.Errorf("did not store user data got %q want %q", store.storeCalls[0], user+data)
		}
	})
}

func getDataRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func newPostSaveRequest(id string, data string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/users/%s/%s", id, data), nil)
	return req
}
