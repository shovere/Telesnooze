import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Data } from "@angular/router";
import { Observable } from "rxjs";

export interface alarms {
    Alarm_ID: number;
    Days: string;
    Time: string;
}


@Injectable({
    providedIn: 'root'
})

export class ApiCallServceService {
    apiUrlRetrieve = 'http://localhost:8123/api/v1/retrieveAlarms"';
    apiUrlDelete = 'http://localhost:8123/api/v1/deleteAlarm';
    apiUrlUpdate = 'http://localhost:8123/api/v1/updateAlarm';


    constructor(private httpClient: HttpClient) { }


    retrieveAlarms(): Observable<alarms[]> {
        return this.httpClient.get<alarms[]>(this.apiUrlRetrieve); 

    }

    deleteAlarm(Alarm_ID: number) {
        return this.httpClient.delete(`${this.apiUrlDelete}/${Alarm_ID}`);
      }

    updateAlarm(Alarm_ID: number) {
        
    }


}

    

    

