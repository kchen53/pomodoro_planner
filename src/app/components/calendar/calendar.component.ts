import { Component, OnInit } from '@angular/core';
import { CalendarView } from 'angular-calendar';
import { Event } from '../calendar/event'
import { EventService } from '../calendar/eventService'
import { CalendarEvent } from 'angular-calendar';
import { MatCalendar } from '@angular/material/datepicker';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
import { v4 as uuidv4 } from 'uuid';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css'],
})

export class CalendarComponent{
  
  viewDate: Date = new Date();
  events: CalendarEvent[] = [];
  event : Event[] = [];

  id = 0;
  name = '';
  start = '2023-04-20T10:00:00';
  end = '2023-04-20T12:00:00';
  color = '';

  constructor(private eventService: EventService) { }

  addEvent() {
    const newEvent: Event = {
      id: this.id,
      name: this.name,
      start: this.start,
      end: this.end,
      color: this.color
    };
  
    this.eventService.addEvent(newEvent).subscribe(
      (event: Event) => {
        console.log('Event added successfully!', event);
        const calendarEvent: CalendarEvent = {
          title: event.name,
          start: new Date(event.start),
          end: new Date(event.end),
          color: {
            primary: event.color,
            secondary: '#FAE3E3'
          }
        };
        this.events.push(calendarEvent);
  
        // Reset the form fields after adding the new event
        this.name = '';
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
