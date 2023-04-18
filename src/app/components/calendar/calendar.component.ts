import { Component, OnInit } from '@angular/core';
import { CalendarView } from 'angular-calendar';
import { Event } from '../calendar/event'
import { EventService } from '../calendar/eventService'

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css'],
})

export class CalendarComponent implements OnInit{

  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;

  setView(view: CalendarView) {
    this.view = view;
  }

  //events
  newEvent: string = '';
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

  addEvents(name : string): void {
    name = name.trim();
    if (!name) { return; }
    const event: Event = { name } as Event;
    this.eventService.addEvent(event).subscribe(newEvent => {
        this.events.push(newEvent);
      });
      this.newEvent = '';
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
