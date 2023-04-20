import { Component, ViewChild, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { ApiCallService, alarms } from './api-call-servce.service';
import { HttpClient } from '@angular/common/http';
import { MatTable } from '@angular/material/table';
import { Observable } from 'rxjs';
import { MatSnackBar } from '@angular/material/snack-bar';


@Component({
  selector: 'app-useralarm',
  styleUrls: ['useralarms.component.css'],
  templateUrl: 'useralarms.component.html',
})
export class TableFilteringExample implements OnInit {

  listalarms: alarms[] = [];
  apiUrlDelete = 'http://localhost:8123/api/v1/deleteAlarm';
  apiUrlUpdate = 'http://localhost:8123/api/v1/updateAlarm';
  showImage = false;
  editElementIndex = -1;
  dataSource = new MatTableDataSource<alarms>();
  isRowBeingEdited = false;
  columnsToDisplay: string[] = [
    'Alarm_ID',
    'days.sunday',
    'days.monday',
    'days.tuesday',
    'days.wednesday',
    'days.thursday',
    'days.friday',
    'days.saturday',
    'time',
    'delete',
    'edit'
  ];

  constructor(
    private apicallservice: ApiCallService,
    private http: HttpClient,
    private snackBar: MatSnackBar
  ) { }

  ngOnInit() {
    this.fetchAlarms();
  }

  @ViewChild(MatTable) myTable!: MatTable<alarms>;

  startEditingRow(element: alarms) {
    this.editElementIndex = element.Alarm_ID;
    this.isRowBeingEdited = true;
  }

  fetchAlarms() {
    this.apicallservice.retrieveAlarms().subscribe((data: any) => {
      this.listalarms = data.alarms;
      this.dataSource.data = this.listalarms;
      this.myTable.renderRows();
      console.log('list of alarms', this.listalarms);
    });
  }

  deleteAlarm(Alarm_ID: number) {
    const requestBody = {
      'Alarm_ID': Alarm_ID
    };
    this.http.post(this.apiUrlDelete, requestBody).subscribe(
      () => {
        this.listalarms = this.listalarms.filter((element) => element.Alarm_ID !== Alarm_ID);
        this.dataSource.data = this.listalarms;
        this.myTable.renderRows();
      },
      (error) => {
        console.error(error);
        alert('Failed to delete the element.');
      }
    );
  }

  updateAlarm(alarms: alarms) {
    const requestBody = {
      'Alarm_ID': alarms.Alarm_ID,
      'time': alarms.time,
      'sunday': alarms.sunday,
      'monday': alarms.monday,
      'tuesday': alarms.tuesday,
      'wednesday': alarms.wednesday,
      'thursday': alarms.thursday,
      'friday': alarms.friday,
      'saturday': alarms.saturday,
      'user_id': alarms.user_id
    };
    this.http.post(this.apiUrlUpdate, requestBody).subscribe(
      (data) => {
        this.isRowBeingEdited = false; // Stop inline text editing
        console.log(data);
        // add any additional code to handle a successful response
      },
      (error) => {
        console.error(error);
        // add any additional code to handle an error response
      }
    );
  }
}
