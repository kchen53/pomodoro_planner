import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';


@Injectable({
  providedIn: 'root'
})
export class LoginService {
  constructor(
      private http: HttpClient
    ) {}

  private loginUrl = 'http://127.0.0.1:8081/login';
  private signUpUrl = 'http://127.0.0.1:8081/signup';

  login(username: string, password: string) {
    return this.http.put(this.loginUrl, { username: username, password: password }, { responseType: 'text' });
  }

  signup(username: string, password: string) {
    return this.http.post(this.signUpUrl, { username: username, password: password }, { responseType: 'text' });
  }
}
