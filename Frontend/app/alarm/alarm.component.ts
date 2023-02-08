import {Component, OnInit} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {HttpClient} from "@angular/common/http";


@Component({
  selector: 'app-alarm',
  templateUrl: './alarm.component.html',
  styleUrls: ['./alarm.component.css']
})
export class AlarmComponent {
    activealarm = this._formBuilder.group({
      sunday : false,
      monday : false,
      tuesday : false,
      wednesday : false,
      thursday : false,
      friday : false,
      saturday : false
    });
  
    constructor(private _formBuilder: FormBuilder) {}
  }



