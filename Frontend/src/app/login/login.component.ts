import { Component } from '@angular/core';
import { Router } from '@angular/router';

/**
 * @title login demo
 */
@Component({
  selector: 'app-login',
  styleUrls: ['./login.component.css'],
  templateUrl: './login.component.html',
})
export class LoginComponent {
  constructor(private router: Router) {}
  username: string = '';
  password: string = '';
  show: boolean = false;
  isLoading = false;
  isSuccess = false;
  isError = false;
  

  submit() {
      fetch('http://localhost:8123/api/v1/login', {
      headers: {
        'content-type': ' application/json',
      },
      method: 'POST',
      body: JSON.stringify({
        username: this.username,
        password: this.password,
      }),
    })
    .then((res) => {
    
      if(res.status != 200){
        this.isLoading = false;
        this.isError = true;
        console.log("FAILED LOGIN");
      }
      else{
        console.log(res);
        console.log("SUCCESSFUL LOGIN");
        this.router.navigate([{ outlets: { home: ['home'] } }]);
        this.isLoading = false;
        this.isSuccess = true;

      }
      
     
    })
    .catch((err) => {
      this.isLoading = false;
      this.isError = true;
      console.log(err);
      console.log("FAILED LOGIN");
    
    });

    this.clear();
  }
  clear() {
    this.username = '';
    this.password = '';
    this.show = true;
  }
}
