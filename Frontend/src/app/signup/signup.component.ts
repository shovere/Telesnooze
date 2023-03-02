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
  phone: string = '';
  email: string = '';
  username: string = '';
  password: string = '';
  passwordconfirm: string = '';
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
  clear() {
    this.username = '';
    this.password = '';
    this.passwordconfirm = '';
    this.phone = '';
    this.show = true;
  }
}
