import { Component, ViewChild,} from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { ApiCallServceService } from './api-call-servce.service';
import { MatTable } from '@angular/material/table';
import { alarms } from './api-call-servce.service';
import { HttpClient } from '@angular/common/http';



const ELEMENT_DATA: alarms[] = [];


@Component({
  selector: 'app-useralarm',
  styleUrls: ['useralarms.component.css'],
  templateUrl: 'useralarms.component.html',
})

export class TableFilteringExample {


    columnsToDisplay: string[] = ['Alarm_ID', 'Days','Time',];
    apiUrlDelete = 'http://localhost:8123/api/v1/deleteAlarm'
  
  

  
  #title = 'mouse-hover';
  showImage: boolean; 
  editElementIndex = -1;

  editCache: { [key: string]: any } = {};
  listOfData: any[] = [];
 
  dataSource = ELEMENT_DATA; 
  

  @ViewChild(MatTable) myTable!: MatTable<alarms>;

  constructor(private apicallservice: ApiCallServceService, 
              private http: HttpClient, ) {

    
    this.showImage = false; 

    this.apicallservice.retrieveAlarms().subscribe(x => {
      this.dataSource = x;
      console.log(this.dataSource)
    })

  } 

  removeItem(Alarm_ID: number): void {
    const url = `${this.apiUrlDelete}/${Alarm_ID}`;
    this.http.delete(url).subscribe(response => {
      console.log('sucess')
    }, error => {
      console.log(error)
    });
  }


  /*
  deleteAlarm(id: number) {
    const url = `http://localhost:8123/api/v1/deleteAlarm/${id}`;
    return this.http.delete(url);
  }


  delete(row: any): void {
    const url =  = this.dataSource.indexOf(row, 0);
    if (index > -1) {
      this.dataSource.splice(index, 1);
    }
    this.myTable.renderRows();
  }*/





  


}





