package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)
All Empty Fields, One Empty Field, Invalid Password, Invalid Phone Number - 10 Digits, Invalid Phone Number - Digits Only, and Optimal Account Creation
func TestCreateUser(t *testing.T) {
	t.Run("All Empty Fields", func(t *testing.T) {
		jsonBody := []byte(`{"email": "", "username": "", "password": "", "phone": ""}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createUser", bodyReader)
		response := httptest.NewRecorder()
		app.createUser(response, request)
		got := response.Body.String();
		want := "Problem: All fields must be filled"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("One Empty Field", func(t *testing.T) {
		jsonBody := []byte(`{"email": "test", "username": "for", "password": "empty", "phone": ""}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createUser", bodyReader)
		response := httptest.NewRecorder()
		app.createUser(response, request)
		got := response.Body.String();
		want := "Problem: All fields must be filled"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Invalid Password", func(t *testing.T) {
		jsonBody := []byte(`{"email": "only", "username": "ascii", "password": "£", "phone": "9999999999"}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createUser", bodyReader)
		response := httptest.NewRecorder()
		app.createUser(response, request)
		got := response.Body.String();
		want := "Problem: Password must only contain ASCII characters"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Invalid Phone Number - 10 digits", func(t *testing.T) {
		jsonBody := []byte(`{"email": "Numbers", "username": "Only", "password": "Ten",  "phone": "99999999999"}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createUser", bodyReader)
		response := httptest.NewRecorder()
		app.createUser(response, request)
		got := response.Body.String();
		want := "Problem: Phone number is invalid - must be 10 digits and only contain numbers"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Invalid Phone Number - Digits Only", func(t *testing.T) {
		jsonBody := []byte(`{"email": "Numbers", "username": "Only", "password": "Ten",  "phone": "aaaaaaaaaa"}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createUser", bodyReader)
		response := httptest.NewRecorder()
		app.createUser(response, request)
		got := response.Body.String();
		want := "Problem: Phone number is invalid - must be 10 digits and only contain numbers"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Optimal Account Creation", func(t *testing.T) {
		jsonBody := []byte(`{"email": "sean.p.hernandez@gmail.com", "username": "Sean", "password": "Hernandez", "phone": "9999999999"}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createUser", bodyReader)
		response := httptest.NewRecorder()
		app.createUser(response, request)
		got := response.Body.String();
		want := "Success"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
}

func TestCreateAlarm(t *testing.T){
	t.Run("No Data",func(t *testing.T) {
		jsonBody := []byte(``)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.createAlarm(response, request)
		got := response.Body.String();
		want :=  "{\"error\":\"Invalid request payload\"}"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
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
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.createAlarm(response, request)
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
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.createAlarm(response, request)
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
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/createAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.createAlarm(response, request)
		got := response.Body.String();
		want := "Success"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
}
