package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_index(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
