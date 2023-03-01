import { LoginComponent } from "src/app/login/login.component";

describe('login.cy.ts', () => {
  it('can mount', () => {
    cy.mount(LoginComponent);
  })


  it('properly renders input fields', () => {
    cy.mount(LoginComponent);
    cy.get('button').should('exist');
    cy.get('mat-form-field.user-prompt').should('exist');
    cy.get('mat-form-field.password-prompt').should('exist');
  })

  it('accepts text inputs', () => {
    cy.mount(LoginComponent);
    const username = 'Allie Gator';
    const password = 'swamplover';

    cy.get('.user-prompt > input').type(username);
    cy.get('.password-prompt > input').type(password);
    
  })

})