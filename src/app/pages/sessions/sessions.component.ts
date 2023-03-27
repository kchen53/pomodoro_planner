import { Component } from '@angular/core';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-sessions',
  templateUrl: './sessions.component.html',
  styleUrls: ['./sessions.component.css'],
  providers: [DatePipe]
})
export class SessionsComponent {
  currentDate: Date | undefined;

  constructor(private datePipe: DatePipe) {}

  ngOnInit() {
    this.currentDate = new Date(); // Set initial date value
    setInterval(() => {
      this.currentDate = new Date(); // Update date every second
    }, 1000);
  }

  handleTimerFinished() {
    console.log('Timer finished!');
    // Perform any necessary actions here
  }
}
