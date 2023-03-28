import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';


@Injectable({
  providedIn: 'root'
})
export class LoginService {
  constructor(
      private http: HttpClient
    ) {}

  private loginUrl = 'http://localhost:8081/login';
  private signUpUrl = 'http://localhost:8081/signup';

  login(username: string, password: string) {
    return this.http.post(this.loginUrl, { username: username, password: password });
  }

  signup(username: string, password: string) {
    return this.http.post(this.signUpUrl, { username: username, password: password });
  }
}
