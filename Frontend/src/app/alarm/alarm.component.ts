import { Component, OnInit } from '@angular/core';
import { Form, FormBuilder } from '@angular/forms';
import { FormGroup, FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

//import { Alarm } from '../alarm';
//import { AlarmService } from '../alarm.service';

@Component({
  selector: 'app-alarm',
  templateUrl: './alarm.component.html',
  styleUrls: ['./alarm.component.css'],
})
export class AlarmComponent {
  constructor(
    //private alarmService: AlarmService,
    private _formBuilder: FormBuilder,
    private http: HttpClient
  ) {}

  time: string = '';

  week = this._formBuilder.group({
    sunday: false,
    monday: false,
    tuesday: false,
    wednesday: false,
    thursday: false,
    friday: false,
    saturday: false,
  });

  days: string = '';

  range = new FormGroup({
    start: new FormControl<Date | null>(null),
    end: new FormControl<Date | null>(null),
  });

  submit() {
    console.log(this.time);

    let tmpDate = new Date();
    let hours = 0;
    let minutes = 0;
    if (this.time.charAt(1) == ':') {
      hours = parseInt(this.time.slice(0, 1));
      minutes = parseInt(this.time.slice(2, 4));
    } else {
      hours = parseInt(this.time.slice(0, 2));
      minutes = parseInt(this.time.slice(3, 5));
    }
    if (this.time.slice(-2) == 'PM') {
      hours += 12;
    }
    tmpDate.setHours(hours);
    tmpDate.setMinutes(minutes);

    console.log(this.week);
    console.log('reach');
    fetch('http://localhost:8123/api/v1/createAlarm', {
      headers: {
        'content-type': ' application/json',
      },
      method: 'POST',
      body: JSON.stringify({
        time: tmpDate.toISOString(),
        days: this.week.value,
      }),
    })
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      });
  }

  /*
  add(time: string): void {
    time= time.trim();
    if (!time) { return; }
    this.alarmService.addAlarm({ time } as Alarm)
      .subscribe(alarm => {
        this.alarms.push(alarm);
      });
  }





  /*
    alarms: Alarm[] = [];




    ngOnInit(): void {
      this.getAlarms();
    }

    getAlarms(): void {
      this.alarmService.getAlarm()
        .subscribe(alarms => this.alarms = alarms);
    }

    add(time: string): void {
      time = time.trim();
      if (!time) {
        return;
      }
      this.alarmService.addAlarm({time} as Alarm)
        .subscribe(time => {
          this.alarms.push(time);
        });
    }
  }

   */
}
