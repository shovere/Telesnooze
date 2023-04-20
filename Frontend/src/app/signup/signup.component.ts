import { Component } from '@angular/core';
import {FormControl, FormGroup, ValidationErrors, ValidatorFn, Validators} from '@angular/forms';
import { passwordMatchValidator} from './signup.validator';
import { HttpClient } from '@angular/common/http';



@Component({
  selector: 'app-signup',
  styleUrls: ['./signup.component.css'],
  templateUrl: './signup.component.html',
})
export class SignupComponent {
 
  isLoading = false;
  isSuccess = false;
  isError = false;

  
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
  
  submit(){
    this.isLoading = true;
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
        if(res.status != 200){
          this.isLoading = false;
          this.isError = true;
          console.log("FAILED SIGNUP");
        }
        else{
          console.log(res);
          console.log("SUCCESSFUL SIGNUP");
          this.isLoading = false;
          this.isSuccess = true;
  
        }
        console.log(res);
       
      })
      .catch((err) => {
        this.isLoading = false;
        this.isError = true;
        console.log("FAILED SIGNUP - FATAL ERROR");
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
     
    this.isLoading = false;
    this.isSuccess = false;
    this.isError = false;
  }
}



