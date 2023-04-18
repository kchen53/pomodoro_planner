import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { LoginService } from './login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  username: string = '';
  password: string = '';
  mode: string = 'Login';
  modeButton: string = 'Sign Up';
  login: boolean = true;
  success: boolean = false;
  showFailed: boolean = false;
  failMessage: string = 'Username or password is incorrect';


  constructor(
    private loginService: LoginService,
    public dialogRef: MatDialogRef<LoginComponent>
  ) {  }

  onSubmit() {
    if (this.login) {
      // login
      this.loginService.login(this.username, this.password).subscribe(
        (data) => {
          if (data.valueOf() == 'true') {
            this.dialogRef.close();
            location.reload();
          }
          else {
            this.password = '';
            this.showFailed = true;
          }
        });
    }
    else {
      // signup
      this.loginService.signup(this.username, this.password).subscribe(
        (data) => {
          if (data.valueOf() == 'true') {
            this.dialogRef.close();
            location.reload();
          }
          else {
            this.password = '';
            this.showFailed = true;
          }
        });
    }

    if (this.mode == 'Login') {
      this.failMessage = 'Username or password is incorrect';
    }
    else {
      this.failMessage = 'Username already exists';
    }

  }

  changeMode() {
    this.login = !this.login;
    this.mode = this.login ? 'Login' : 'Sign Up';
    this.modeButton = this.login ? 'Sign Up' : 'Login';
  }
}
