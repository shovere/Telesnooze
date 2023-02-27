package main

import "testing"

func TestSetAlarm(t *testing.T){
	var a alarm;
	app := &App{}
	app.initializeApp()
	got, err:= app.writeDBAlarm(a);
}

