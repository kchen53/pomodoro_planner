import { Component } from '@angular/core';
import { MatDialog} from '@angular/material/dialog';
import { LoginComponent } from '../login/login.component';
import { NavbarService } from './navbar.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  username: string;
  loggedIn: boolean;

  constructor(
    public dialog: MatDialog,
    private navbarService: NavbarService
    ) {
      this.username = 'Anonymous';
      this.loggedIn = false;
    }

  ngOnInit() {
    this.navbarService.getUsername().subscribe(
      (data) => {
        if (data.valueOf() != '"Anonymous"') {
          this.username = data.valueOf();
          this.username = this.username.substring(1, this.username.length - 1);
          this.loggedIn = true;
        }
      });
  }

  openDialog(){
    this.dialog.open(LoginComponent);
  }

  logout() {
    this.navbarService.logout().subscribe(
      (data) => {
        console.log(data);
        this.loggedIn = false;
      });
  }

}
