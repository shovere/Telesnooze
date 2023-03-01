import { HomeComponent } from "src/app/home/home.component";

describe('signup.cy.ts', () => {
  it('can mount', () => {
    cy.mount(HomeComponent);
  })


  it('welcomes the user', () => {
    cy.mount(HomeComponent);
    cy.get('.home-header');
  
  })
  it('has all three tabs present', () => {
    cy.mount(HomeComponent);
    cy.get('[label="Your Stats"]');
    cy.get('[label="Configure Dates"]');
    cy.get('[label="Customize Questions"]');
   
  })
})