import { Component, ViewChild,} from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { MatTable } from '@angular/material/table';
import { HttpClient } from '@angular/common/http';


export interface Alarm {
  Alarm_ID: number;
  Days: string;
  Time: string;
}

const ELEMENT_DATA: Alarm[] = [
  {Alarm_ID: 1, Days: 'Tuesday', Time: '08:00'},
  {Alarm_ID: 2, Days: 'Wendesday', Time: '08:00'},
  {Alarm_ID: 3, Days: 'Thursday', Time: '08:00'},
  {Alarm_ID: 4, Days: 'Friday', Time: '08:00'},
];

/**
 * @title Basic use of `<table mat-table>`
 */
@Component({
  selector: 'table-basic-example',
  styleUrls: ['test.component.css'],
  templateUrl: 'test.component.html',
})
export class TestComponent {
  displayedColumns: string[] = ['Alarm_ID', 'Days', 'Time','edit', 'delete'];
  dataSource = ELEMENT_DATA;

  #title = 'mouse-hover';

  editElementIndex = -1;

  editCache: { [key: string]: any } = {};
  listOfData: any[] = [];
 
  

  @ViewChild(MatTable) myTable!: MatTable<Alarm>;

  delete(row: any): void {
    const index = this.dataSource.indexOf(row, 0);
    if (index > -1) {
      this.dataSource.splice(index, 1);
    }
    this.myTable.renderRows();
  }
}


/**  Copyright 2023 Google LLC. All Rights Reserved.
    Use of this source code is governed by an MIT-style license that
    can be found in the LICENSE file at https://angular.io/license */