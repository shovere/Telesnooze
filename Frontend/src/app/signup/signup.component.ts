import { Component } from '@angular/core';
/**
* @title login demo
*/
@Component({
  selector: 'app-signup',
  styleUrls: ['./signup.component.css'],
  templateUrl: './signup.component.html',
})
export class SignupComponent {
  phone : string = "";
  email: string = "";
  username: string = "";
  password: string = "";
  passwordconfirm: string = "";
  show: boolean = false;
  submit() {
    console.log("user name is " + this.username)
    this.clear();
  }
  clear() {
    this.username = "";
    this.password = "";
    this.passwordconfirm = "";
    this.phone = "";
    this.show = true;
  }
}
