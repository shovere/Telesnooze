import {Component, OnInit} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'checkbox-reactive-forms-example',
  templateUrl: 'alarm.html',
  styleUrls: ['alarm.css'],
})

export class activealarmgroup {
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

  export class AppMdodule {
  postData = {
    test: 'mycontent',
  }
  url: 'foo.com'

  constructor(private http: HttpClient) {
    this.http.post(this.url, this.postData).toPromise().then(data => {
      console.log(data);
    });
  }
  }