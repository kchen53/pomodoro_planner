import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';
import { Event } from '../calendar/event'

@Injectable({
  providedIn: 'root'
})

export class EventService {

  private apiUrl = 'http://127.0.0.1:8081/event';
  //private apiUrl = 'https://jsonplaceholder.typicode.com/todos';

  constructor(private http: HttpClient) { }

  getEvents(): Observable<Event[]> {
    this.http.delete(`http://127.0.0.1:8081/event/22`);
    return this.http.get<Event[]>(this.apiUrl);
  }

  addEvent(event: Event): Observable<Event> {

    return this.http.post<Event>(this.apiUrl, event); 
  }

  updateEvent(event: Event): Observable<Event> {
    const url = `${this.apiUrl}/${event.id}`;
    return this.http.put<Event>(url, event);
  }

  deleteEvent(id: number): Observable<{}> {
    const url = `${this.apiUrl}/${id}`;
    console.log("deleting ", url);
    return this.http.delete(url);
  }
}
