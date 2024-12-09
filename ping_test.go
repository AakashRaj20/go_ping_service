package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	//call the ping function
	ping(server.URL)
}