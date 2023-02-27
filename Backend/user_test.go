package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetAlarm(t *testing.T){
	t.Run("Bad Time", func(t *testing.T) {
		jsonBody := []byte(`{"time": "202-27T17:43:35.668Z",
							 "days": {
								"sunday": false, 
								"monday": false, 
								"tuesday": true, 
								"wednesday": false, 
								"thursday": false, 
								"friday": false, 
								"saturday": false
							 }}`)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/setAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.setAlarm(response, request)
		got := response.Body.String();
		want := "Timestamp is not in ISO format"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Bad Weekdays",func(t *testing.T) {
		jsonBody := []byte(`{"time": "2023-02-27T17:43:35.668Z",
							 "days": {
								"sunday": false, 
								"monday": false, 
								"tuesday": false, 
								"wednesday": false, 
								"thursday": false, 
								"friday": false, 
								"saturday": false
							 }}`)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/setAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.setAlarm(response, request)
		got := response.Body.String();
		want := "Problem: Week needs at least one true value OR JSON be malformed"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Optimal Test Alarm", func(t *testing.T) {
		jsonBody := []byte(`{"time": "2023-02-27T17:43:35.668Z",
							 "days": {
								"sunday": false, 
								"monday": false, 
								"tuesday": true, 
								"wednesday": false, 
								"thursday": false, 
								"friday": false, 
								"saturday": false
							 }}`)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/setAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.setAlarm(response, request)
		got := response.Body.String();
		want := "Success"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
}

