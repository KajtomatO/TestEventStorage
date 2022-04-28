package server_lib

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryDataStore()
	server := DataServer{store}
	user := "Pepper"
	data := "d1"

	server.ServeHTTP(httptest.NewRecorder(), newPostSaveRequest(user, data))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, getDataRequest(user))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), data)
}
