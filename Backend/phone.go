package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
);

func callNumber() {
	err := godotenv.Load();
	
	if(err != nil) {
		fmt.Print("Error loading .env file")
	}

    accountSid := os.Getenv("ACCOUNT_SID")
    authToken := os.Getenv("AUTH_TOKEN")
    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })
	from := "+18884956375"
	to := "+16035689902"

	params := &twilioApi.CreateCallParams{}
    params.SetTo(to)
    params.SetFrom(from)
    params.SetUrl("http://demo.twilio.com/docs/voice.xml")

    resp, err := client.Api.CreateCall(params)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Call Status: " + *resp.Status)
        fmt.Println("Call Sid: " + *resp.Sid)
        fmt.Println("Call Direction: " + *resp.Direction)
    }
}