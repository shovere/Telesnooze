import { SignupComponent } from "../src/app/signup/signup.component";

describe('signup.cy.ts', () => {
  it('can mount', () => {
    cy.mount(SignupComponent);
  })


  it('properly renders input fields', () => {
    cy.mount(SignupComponent);
    cy.get('button').should('exist');
    cy.get('mat-form-field.email-prompt').should('exist');
    cy.get('mat-form-field.user-prompt').should('exist');
    cy.get('mat-form-field.password-prompt').should('exist');
    cy.get('mat-form-field.password-confirm-prompt').should('exist');
    cy.get('mat-form-field.phone-prompt').should('exist');
  })

  it('accepts text inputs', () => {
    cy.mount(SignupComponent);
    const username = 'Allie Gator';
    const email = 'ag@ufl.edu';
    const password = 'swamplover';
    const phone = '123456789';


    cy.get('.user-prompt > input').type(username);
    cy.get('.password-prompt > input').type(password);
    cy.get('.password-confirm-prompt > input').type(password);
    cy.get('.email-prompt > input').type(email);
    cy.get('.phone-prompt > input').type(phone);
  })


  it('prevents invalid email submissions', () => {
    cy.mount(SignupComponent);
    const username = 'Allie Gator';
    const email = 'agufl.edu';
    const password = 'swamplover';
    const phone = '0123456789';
    cy.get('.user-prompt > input').type(username);
    cy.get('.password-prompt > input').type(password);
    cy.get('.password-confirm-prompt > input').type(password);
    cy.get('.email-prompt > input').type(email);
    cy.get('.phone-prompt > input').type(phone);
    cy.get('button').should('be.disabled');
  })

  it('prevents invalid phone number submissions', () => {
    cy.mount(SignupComponent);
    const username = 'Allie Gator';
    const email = 'agufl.edu';
    const password = 'swamplover';
    const phone = 'abcdefghasduihas';
    cy.get('.user-prompt > input').type(username);
    cy.get('.password-prompt > input').type(password);
    cy.get('.password-confirm-prompt > input').type(password);
    cy.get('.email-prompt > input').type(email);
    cy.get('.phone-prompt > input').type(phone);
    cy.get('button').should('be.disabled');
  })


  it('prevents non-matching passwords', () => {
    cy.mount(SignupComponent);
    const username = 'Allie Gator';
    const email = 'agufl.edu';
    const password = 'swamplover';
    const confirmPassword = 'asiodjaosidjaosijdoasijd';
    const phone = '1234567890';
    cy.get('.user-prompt > input').type(username);
    cy.get('.password-prompt > input').type(password);
    cy.get('.password-confirm-prompt > input').type(confirmPassword);
    cy.get('.email-prompt > input').type(email);
    cy.get('.phone-prompt > input').type(phone);
    cy.get('button').should('be.disabled');
  })



  /*it('accepts button input', () => {
    cy.mount(SignupComponent);
    cy.get('.signup-button').click();

  })  //TODO: find a way to pass this test without switching to a different url
  */

})