import { Component, EventEmitter, Output } from '@angular/core';
import { interval, Subscription } from 'rxjs';

@Component({
  selector: 'app-timer',
  templateUrl: './timer.component.html',
  styleUrls: ['./timer.component.css']
})
export class TimerComponent {
  interval: any;

public time: { minutes: number, seconds: number } = { minutes: 0, seconds: 0 };
public originalTime: { minutes: number, seconds: number }= { minutes: 0, seconds: 0 };
public timeString: string = '';
public isTimerStarted: boolean = false;
@Output() timerFinished = new EventEmitter<void>();

constructor() {
  this.timeString = ''; // Set default time to 2 minutes and 30 seconds
}

startTimer() {
  this.isTimerStarted = true;
  // Store the original value of the timer before starting it
  this.originalTime = { minutes: this.time.minutes, seconds: this.time.seconds };

  // Start the timer
  this.interval = setInterval(() => {
    if (this.time.minutes === 0 && this.time.seconds === 0) {
      this.stopTimer();
      this.timerFinished.emit();
    } else if (this.time.seconds === 0) {
      this.time.minutes--;
      this.time.seconds = 59;
    } else {
      this.time.seconds--;
    }
  }, 1000);

  //this.timeString = "";
}

  stopTimer() {
    //this.isTimerStarted = false;
    clearInterval(this.interval);
  }

  resetTimer() {
    //this.isTimerStarted = false;
    // Reset the time object to the original value
    this.time = { minutes: this.originalTime.minutes, seconds: this.originalTime.seconds };
  }

  clearTimer(){
    this.isTimerStarted = false;
    clearInterval(this.interval);
    this.timeString = '';
    this.time = { minutes: 0, seconds: 0 };
  }

  getTimer() {
    const minutes = this.time.minutes;
    const seconds = this.time.seconds;

    if(this.isTimerStarted==true){
    return `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    }
    else{
      return `00:00`;
    }
  }

  updateTime() {
    const parts = this.timeString.split(':');
    this.time.minutes = Number(parts[0]);
    this.time.seconds = Number(parts[1]);
  }

}
