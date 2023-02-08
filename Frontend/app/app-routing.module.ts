import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router'; 
import { SignupComponent } from './signup/signup.component';
import { LoginComponent } from './login/login.component';
import { TitleComponent } from './title/title.component';
import { BrowserModule } from '@angular/platform-browser';
import { HomeComponent } from './home/home.component';



const routes: Routes = [
  {
    component: SignupComponent,
    path: 'signup',
    outlet: "signup"
    
  },
  {
    path: '',
    pathMatch : 'full',
    component: LoginComponent,
    outlet: 'login'
  },
  {
    path: 'home',
    outlet: 'home',
    component: HomeComponent,
  }
];
@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    BrowserModule,
    RouterModule.forRoot(routes)
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
