package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticateLogin(t *testing.T) {
	t.Run("Authenticating Login", func(t *testing.T) {
		jsonBody := []byte(`{"username": "Sean", "password": "Hernandez"}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/authenticationEndpoint", bodyReader)
		response := httptest.NewRecorder()
		app.authenticationEndpoint(response, request)
		got := response.Body.String();
		want := "Successful find"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Failed Login", func(t *testing.T) {
		jsonBody := []byte(`{"username": "uvdnioqcwdhnclq", "password": "erjnlirnle"}`)
		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/authenticationEndpoint", bodyReader)
		response := httptest.NewRecorder()
		app.authenticationEndpoint(response, request)
		got := response.Body.String();
		want := "Problem: Username or password is incorrect"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
}

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
		jsonBody := []byte(`{
			"user_id": "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97",
			"time": "202-27T17:43:35.668Z",
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
		jsonBody := []byte(`{
			"user_id": "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97",
			"time": "2023-02-27T17:43:35.668Z",
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
		jsonBody := []byte(`{
							"user_id": "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97",
							"time": "2023-02-27T14:26:35.668Z",
							 "days": {
								"sunday": false, 
								"monday": true, 
								"tuesday": true, 
								"wednesday": true, 
								"thursday": true, 
								"friday": true, 
								"saturday": true
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


func TestCallNumberStandard(t *testing.T){
	t.Run("Standard Run", func(t *testing.T) {
		got := callNumber("+16035689902")
		want := "queued"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	
}

func TestCallNumberFail(t *testing.T){
	t.Run("Unverified number", func(t *testing.T) {
		got := callNumber("+11111111111")
		want := "Could not place call, check logs for error"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	
}

func TestRetreiveAlarms(t *testing.T){
	t.Run("Basic Test",func(t *testing.T) {
		jsonBody := []byte(`{"user_id": "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97"}`)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/retrieveAlarms", bodyReader)
		response := httptest.NewRecorder()
	
		app.retrieveAlarms(response, request)
		
		var tmpRetAlarm retAlarms
		decoder := json.NewDecoder(response.Body)
	
		errDecode := decoder.Decode(&tmpRetAlarm)
		fmt.Printf("%v", tmpRetAlarm.User_ID)
		if errDecode != nil {
			fmt.Println(errDecode)
			return
		}

		fmt.Println("array")
		fmt.Println(tmpRetAlarm)
		got := len(tmpRetAlarm.Alarms);
		if got == 0 {
			t.Errorf("response body is wrong, got %q", got)
		}
	})

	//should not return any alarms
	t.Run("No user id",func(t *testing.T) {
		jsonBody := []byte(`{"user_id": ""}`)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/retrieveAlarms", bodyReader)
		response := httptest.NewRecorder()
	
		app.retrieveAlarms(response, request)
		
		var tmpRetAlarm retAlarms
		decoder := json.NewDecoder(response.Body)
	
		errDecode := decoder.Decode(&tmpRetAlarm)
		fmt.Printf("%v", tmpRetAlarm.User_ID)
		if errDecode != nil {
			fmt.Println(errDecode)
			return
		}

		fmt.Println("array")
		fmt.Println(tmpRetAlarm)
		got := len(tmpRetAlarm.Alarms);
		if got != 0 {
			t.Errorf("response body is wrong, got %q", got)
		}
	})
}

func TestUpdateAlarm(t *testing.T){
	t.Run("No Data",func(t *testing.T) {
		jsonBody := []byte(``)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/updateAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.updateAlarm(response, request)
		got := response.Body.String();
		want :=  "{\"error\":\"Invalid request payload\"}"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Bad Time", func(t *testing.T) {
		jsonBody := []byte(`{
			"user_id": "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97",
			"alarm_id": "1f82b94c-3d6e-4cef-8814-74ed20df8fd2",
			"time": "202-27T17:43:35.668Z",
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
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/updateAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.updateAlarm(response, request)
		got := response.Body.String();
		want := "Timestamp is not in ISO format"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Bad Weekdays",func(t *testing.T) {
		jsonBody := []byte(`{
			"user_id": "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97",
			"alarm_id":  "1f82b94c-3d6e-4cef-8814-74ed20df8fd2",
			"time": "2023-02-27T17:15:35.668Z",
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
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/updateAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.updateAlarm(response, request)
		got := response.Body.String();
		want := "Problem: Week needs at least one true value OR JSON be malformed"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
	t.Run("Optimal Test Alarm", func(t *testing.T) {
		jsonBody := []byte(`{
							"user_id": "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97",
							"alarm_id": "1f82b94c-3d6e-4cef-8814-74ed20df8fd2",
							"time":"2023-02-27T17:18:35.668Z",
							 "days": {
								"sunday": true, 
								"monday": true, 
								"tuesday": true, 
								"wednesday": true, 
								"thursday": true, 
								"friday": false, 
								"saturday": false
							 }}`)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/updateAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.updateAlarm(response, request)
		got := response.Body.String();
		want := "Success"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
}

func TestDeleteAlarm(t *testing.T){
	t.Run("No Data",func(t *testing.T) {
		jsonBody := []byte(``)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/deleteAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.deleteAlarm(response, request)
		got := response.Body.String();
		want :=  "{\"error\":\"Invalid request payload\"}"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
		
	t.Run("Delete Alarm", func(t *testing.T) {
		jsonBody := []byte(`{"alarm_id": "0664c23d-673c-47c4-85d6-97e77203f877"}`)
 		bodyReader := bytes.NewReader(jsonBody)
		app := &App{}
		app.initializeApp()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/deleteAlarm", bodyReader)
		response := httptest.NewRecorder()
		app.deleteAlarm(response, request)
		got := response.Body.String();
		want := "Success"
		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})
}