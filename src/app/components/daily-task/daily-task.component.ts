import { Component } from '@angular/core';
import { CalendarComponent } from '../calendar/calendar.component';
import { CalendarEvent, CalendarView } from 'angular-calendar';
import { Event } from '../calendar/event'
import { EventService } from '../calendar/eventService'
import { Subject } from 'rxjs';

@Component({
  selector: 'app-daily-task',
  templateUrl: './daily-task.component.html',
  styleUrls: ['./daily-task.component.css']
})
export class DailyTaskComponent {
  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;

  setView(view: CalendarView) {
    this.view = view;
  }

  events: CalendarEvent[] = [];

  title = '';
  start = '';
  end = '';
  color = "#ad2121";

  constructor(private eventService: EventService) { }

  ngOnInit() {
    this.getEvents();
  }

  getEvents() {
    this.eventService.getEvents().subscribe(
      (events: Event[]) => {
        console.log('Events retrieved successfully!', events);
        this.events = events.map(event => ({
          title: event.title,
          start: new Date(event.start),
          end: new Date(event.end),
          color: {
            primary: "#ad2121",
            secondary: '#FAE3E3'
          }
        }));
      },
      (error) => {
        console.log('Error retrieving events:', error);
      }
    );
  }
}
