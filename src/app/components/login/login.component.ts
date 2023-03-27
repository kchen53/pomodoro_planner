import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username: string = '';
  password: string = '';

  constructor(
    public dialogRef: MatDialogRef<LoginComponent>
  ) {  }

  login() {
    console.log('Username: ' + this.username);
  }

  onSubmit() {
    this.dialogRef.close();
  }
}
