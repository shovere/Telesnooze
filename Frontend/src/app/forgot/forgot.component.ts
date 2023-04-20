import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-forgot',
  templateUrl: './forgot.component.html',
  styleUrls: ['./forgot.component.css']
})
export class ForgotComponent {
  username: string = '';
  password: string = '';
  phone: string = '';
  show: boolean = false;
  isLoading = false;
  isSuccess = false;
  isError = false;


  public forgotForm : FormGroup = new FormGroup({
    'userControl' : new FormControl('',[Validators.required]),
    'passwordControl' : new FormControl('' , [Validators.required]),
    'phoneControl' : new FormControl('',[Validators.required, Validators.pattern("^((\\+91-?)|0)?[0-9]{10}$")] ) 
  });

  submit(){
    this.isLoading = true;
    fetch('http://localhost:8123/api/v1/createUser', {
      headers: {
        'content-type': ' application/json',
      },
      method: 'POST',
      body: JSON.stringify({
        username: this.username,
        password: this.password,
        phone: this.phone,
      }),
    })
      .then((res) => {
        if(res.status != 200){
          this.isLoading = false;
          this.isError = true;
          console.log("FAILED PASSWORD RESET");
        }
        else{
          console.log(res);
          console.log("SUCCESSFUL PASSWORD RESET");
          this.isLoading = false;
          this.isSuccess = true;
  
        }
        console.log(res);
       
      })
      .catch((err) => {
        this.isLoading = false;
        this.isError = true;
        console.log("FAILED PASSWORD RESET - FATAL ERROR");
        console.log(err);
      
      });
      
    this.clear();
  }
  submissionNotReady() : boolean {
    const hasErrors = Object.values(this.forgotForm.controls).some(control => control.errors !== null);
    if(hasErrors){
      return true;
    }
    return false;
  }
  clear() {
    this.username = "";
    this.password = "";
    this.phone = "";
    this.show = true;
     
    this.isLoading = false;
    this.isSuccess = false;
    this.isError = false;
  }

}
