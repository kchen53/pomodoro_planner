import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../environments/environment';
import {map} from 'rxjs/operators';

@Injectable()
export class HelloWorldService {

  constructor(private http: HttpClient) { }

  getTitle() {
    return this.http.get(`${environment.serverUrl}/hello-world`)
      .pipe(map(response => response));
  }

}


