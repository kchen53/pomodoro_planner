import { Component, OnInit } from '@angular/core';
import { CalendarView } from 'angular-calendar';
import { Event } from '../calendar/event'
import { EventService } from '../calendar/eventService'
import { CalendarEvent } from 'angular-calendar';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css'],
})

export class CalendarComponent implements OnInit{

  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;

  startTimePicker = 'start-time-picker';

  setView(view: CalendarView) {
    this.view = view;
  }

  //events
  newEvent: string = '';
  newEventName: string = '';
  newEventDate: string = '';
  newEventStartTime: string = '';
  newEventEndTime: string = '';
  newEventRepeat: number = 0;
  events: Event[] = [];

  constructor(private eventService: EventService) { }

  onEnter() {
    this.newEvent = '';
  }

  ngOnInit() {
    this.getEvents();
  }

  getEvents(): void { 
    this.eventService.getEvents() 
      .subscribe(events => this.events = events); 
  }

  addEvents(name: string, date: string, startTime: string, endTime: string, repeat: number): void {
    name = name.trim();
    if (!name || !date || !startTime || !endTime || !repeat) {
      return;
    }
  
    const event: Event = { name, date, startTime, endTime, repeat } as Event;
    this.eventService.addEvent(event).subscribe(newEvent => {
      this.events.push(newEvent);
    });
    this.newEventName = '';
    this.newEventDate = '';
    this.newEventStartTime = '';
    this.newEventEndTime = '';
    this.newEventRepeat = 0;
  }
  
  deleteEvent(event: Event): void {
    this.events = this.events.filter(e => e !== event);
    this.eventService.deleteEvent(event.id).subscribe();
  } 
  
  // addEvents(task : string): void {
  //   task = task.trim();
  //   if (!task) { return; }
  //   const event: Event = { task } as Event;
  //   this.eventService.addEvent(event).subscribe(newEvent => {
  //       this.events.push(newEvent);
  //     });
  //     this.newEvent = '';
  // }
}
