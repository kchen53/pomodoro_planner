import { Component, OnInit } from '@angular/core';
import { CalendarView } from 'angular-calendar';
import { Event } from '../calendar/event'
import { EventService } from '../calendar/eventService'
import { CalendarEvent } from 'angular-calendar';
import { MatCalendar } from '@angular/material/datepicker';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css'],
})

export class CalendarComponent implements OnInit {
  
  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  events: CalendarEvent[] = [];

  title = '';
  start = '';
  end = '';
  color = '';

  colorOptions = [
    { name: 'Red', value: '#FF0000' },
    { name: 'Green', value: '#00FF00' },
    { name: 'Blue', value: '#0000FF' },
    // add more color options as desired
  ];
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
            primary: event.color,
            secondary: '#FAE3E3'
          }
        }));
      },
      (error) => {
        console.log('Error retrieving events:', error);
      }
    );
  }

  addEvent() {
    const newEvent: Event = {
      id: 0,
      title: this.title,
      start: this.start,
      end: this.end,
      color: this.color
    };
  
    this.eventService.addEvent(newEvent).subscribe(
      (event: Event) => {
        console.log('Event added successfully!', event);
        const calendarEvent: CalendarEvent = {
          title: event.title,
          start: new Date(event.start),
          end: new Date(event.end),
          color: {
            primary: event.color,
            secondary: '#FAE3E3'
          }
        };
        this.events.push(calendarEvent);
  
        // Reset the form fields after adding the new event
        this.title = '';
        this.start = '';
        this.end = '';
        this.color = '';
      },
      (error) => {
        console.log('Error adding event:', error);
      }
    );
  }
}
