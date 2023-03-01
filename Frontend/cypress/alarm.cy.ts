import { AlarmComponent } from "src/app/alarm/alarm.component";
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import {HttpClientModule} from '@angular/common/http';

describe('alarm.cy.ts', () => {

  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule], 
    providers: [AlarmComponent]
  }));

  it('can mount', () => {
    cy.mount(AlarmComponent);
  })

   it('renders date picker', () => {
    cy.mount(AlarmComponent);
    cy.get('section.date-selector').should('exist');
  })

  it('renders days of week checkbox ', () => {
    cy.mount(AlarmComponent);
    cy.get('section.week').should('exist');
  })

  it('renders time picker', () => {
    cy.mount(AlarmComponent);
    cy.get('div.timepicker').should('exist');
  })

  it('renders input field', () => {
    cy.mount(AlarmComponent);
    cy.get('mat-form-field.date-prompt').should('exist'); 
    cy.get('mat-form-field.time-prompt').should('exist');
  })

  it('accepts text inputs', () => {
    cy.mount(AlarmComponent);
    const time = '06:00';
    const days = '2023-03-30T04:00:00.000Z';

    cy.get('.time-prompt > input').type(time);
    cy.get('.date-prompt > input').type(days);
  })

  it('accepts button input', () => {
    cy.mount(AlarmComponent);
    cy.get('.alarm-button').click();
  })  


})
