import { Component } from '@angular/core';
import { CalendarView } from 'angular-calendar';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css'],
})
export class CalendarComponent {
  // selectedDate: Date = new Date();
  // startAt: Date = new Date(new Date().getFullYear(), new Date().getMonth(), 1);

  // constructor() {}

  // onDateChange() {
  //   this.startAt = new Date(
  //     this.selectedDate.getFullYear(),
  //     this.selectedDate.getMonth(),
  //     1
  //   );
  // }

  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;

  setView(view: CalendarView) {
    this.view = view;
  }
}
