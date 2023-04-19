import { Component, EventEmitter, Output } from '@angular/core';
import { interval, Subscription } from 'rxjs';

@Component({
  selector: 'app-timer',
  templateUrl: './timer.component.html',
  styleUrls: ['./timer.component.css']
})
export class TimerComponent {
  interval: any;

public time: {hours:number, minutes: number, seconds: number } = {hours:0, minutes: 5, seconds: 0 };
public isTimerStarted: boolean = false;
public isTimerPaused: boolean = false;
public resettingTimer: boolean = false;
@Output() timerFinished = new EventEmitter<void>();

constructor() {
}

startTimer() {
  if (!(this.time.hours === 0 && this.time.minutes === 0 && this.time.seconds === 0)) {
    this.isTimerStarted = true;

    // Start the timer
    this.interval = setInterval(() => {
      if (this.time.hours === 0 && this.time.minutes === 0 && this.time.seconds <= 1) {
        this.stopTimer();
        this.timerFinished.emit();
      }
      else if (this.time.minutes === 0 && this.time.seconds === 0) {
        this.time.hours--;
        this.time.minutes = 59;
        this.time.seconds = 59;
      }
      else if (this.time.seconds === 0) {
        this.time.minutes--;
        this.time.seconds = 59;
      }
      else {
        this.time.seconds--;
      }
    }, 1000);
  }

}

  stopTimer() {
    if (this.resettingTimer = true) {
      this.resettingTimer = false;
      this.isTimerStarted = false;
      this.isTimerPaused = false;
      clearInterval(this.interval);
    }
    else {
      this.isTimerPaused = true;
    }

  }

  resetTimer() {
    this.stopTimer();
    this.time = {hours: 0, minutes: 5, seconds: 0};
    this.resettingTimer = true;
  }

  getTimer() {
    const hours = this.time.hours;
    const minutes = this.time.minutes;
    const seconds = this.time.seconds;

    if(this.isTimerStarted==true){
    return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    }

    return ``;
  }

}
