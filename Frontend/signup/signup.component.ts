import { Component } from '@angular/core';
import {FormControl, FormGroup, ValidationErrors, ValidatorFn, Validators} from '@angular/forms';
import { passwordMatchValidator} from './signup.validator';


@Component({
  selector: 'app-signup',
  styleUrls: ['./signup.component.css'],
  templateUrl: './signup.component.html',
})
export class SignupComponent {
 
  public signupForm : FormGroup = new FormGroup({
    'userControl' : new FormControl('',[Validators.required]),
    'emailControl' : new FormControl('', [Validators.required, Validators.email]),
    'passwordControl' : new FormControl('' , [Validators.required]),
    'confirmPasswordControl' : new FormControl('',[Validators.required] ),  
    'phoneControl' : new FormControl('',[Validators.required, Validators.pattern("^((\\+91-?)|0)?[0-9]{10}$")] ) 

  }, { validators: passwordMatchValidator });
  phone : string = "";
  email: string = "";
  username: string = "";
  password: string = "";
  passwordconfirm: string = "";
  show: boolean = false;




  


  submit() {
    console.log('user name is ' + this.username);
    fetch('http://localhost:8123/api/v1/createUser', {
      headers: {
        'content-type': ' application/json',
      },
      method: 'POST',
      body: JSON.stringify({
        email: this.email,
        username: this.username,
        password: this.password,
        phone: this.phone,
      }),
    })
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      });
    this.clear();
  }
  checkEmail(){
    const emailControl = this.signupForm.get('emailControl');
    if (emailControl == null){
      return '';
    }
    if(emailControl.hasError('required')) {
      return 'You must enter a value';
    }

    return emailControl.hasError('email') ? 'Not a valid email' : '';
  }

  submissionNotReady() : boolean {
    const hasErrors = Object.values(this.signupForm.controls).some(control => control.errors !== null);
    if(this.signupForm.errors?.['passwordMismatch']){
      return true;
    }
    if(hasErrors){
      return true;
    }
    return false;
  }

  clear() {
    this.username = "";
    this.password = "";
    this.passwordconfirm = "";
    this.phone = "";
    this.show = true;
  }
}

