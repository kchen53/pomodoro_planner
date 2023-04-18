import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class NavbarService {
  constructor(
      private http: HttpClient
    ) {}

  private getUsernameUrl = 'http://127.0.0.1:8081/userdata';
  private logoutUrl = 'http://127.0.0.1:8081/logout';

  getUsername() {
    return this.http.get(this.getUsernameUrl, {responseType: 'text'});
  }

  logout() {
    return this.http.put(this.logoutUrl, {responseType: 'boolean'});
  }
}
