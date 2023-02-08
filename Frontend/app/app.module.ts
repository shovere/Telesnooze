import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

import { MatButton } from '@angular/material/button';
import {MaterialExampleModule} from '../material.module';
import {activealarmgroup} from './alarm';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatNativeDateModule} from '@angular/material/core';
import {HttpClientModule} from '@angular/common/http';
import {NgxMaterialTimepickerModule} from "ngx-material-timepicker";
import { AlarmComponent } from './alarm/alarm.component';

@NgModule({
  declarations: [activealarmgroup, AlarmComponent],
  imports: [
    BrowserAnimationsModule,
    BrowserModule,
    FormsModule,
    HttpClientModule,
    MatNativeDateModule,
    MaterialExampleModule,
    ReactiveFormsModule,
    NgxMaterialTimepickerModule,
  ],
  providers: [],
  bootstrap: [activealarmgroup],
})
export class AppModule {}
