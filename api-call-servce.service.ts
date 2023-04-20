import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Data } from "@angular/router";
import { Observable } from "rxjs";

export interface alarms {
    user_id: number;
    time: string;
    Alarm_ID: number;
      sunday: boolean;
      monday: boolean;
      tuesday: boolean;
      wednesday: boolean;
      thursday: boolean;
      friday: boolean;
      saturday: boolean;
  }
  
  

@Injectable({
    providedIn: 'root'
})
export class ApiCallService {
    apiUrlRetrieve = 'http://localhost:8123/api/v1/retrieveAlarms';
    apiUrlDelete = 'http://localhost:8123/api/v1/deleteAlarm';
    apiUrlUpdate = 'http://localhost:8123/api/v1/updateAlarm';

    constructor(private httpClient: HttpClient) { }

    retrieveAlarms(): Observable<alarms[]> {
        return this.httpClient.post<alarms[]>(this.apiUrlRetrieve, {user_id: "83f18bdf-2e8f-4cd0-bfba-8dd0ec79aa97"});

    }

    deleteAlarm(Alarm_ID: number) {
        return this.httpClient.delete(`${this.apiUrlDelete}/${Alarm_ID}`);

    }




    updateAlarm(alarm: alarms): Observable<alarms> {
        return this.httpClient.post<alarms>(`${this.apiUrlUpdate}/${alarm.Alarm_ID}`, alarm);
      }

    }




