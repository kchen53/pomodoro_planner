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
  username: string = '';
  loggedIn: boolean = false;

  constructor(
    public dialog: MatDialog,
    private navbarService: NavbarService
    ) {}

  ngOnInit() {
    this.navbarService.getUsername().subscribe(
      (data) => {
        this.username = data.valueOf()
      });

    if (this.username != 'Anonymous') {
      this.loggedIn = true;
    }
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
