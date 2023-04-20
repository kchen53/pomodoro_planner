import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from 'src/app/app.module';
import { CalendarComponent } from './calendar.component';
import { EventService } from './eventService';

describe('CalendarComponent', () => {
  let component: CalendarComponent;
  let fixture: ComponentFixture<CalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule,
      ],
      declarations: [ CalendarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CalendarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //New Unit Test
  it('check if the calendar displays the correct month', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('h2').textContent).toContain('April 2023');
  });

  it('check if getEvents function was called', () => {
    expect(component.getEvents).toHaveBeenCalled;
  });

  it('check if addEvent function was called', () => {
    expect(component.addEvent).toHaveBeenCalled;
  });

  it('check if deleteEvent function was called', () => {
    expect(component.deleteEvent).toHaveBeenCalled;
  });
  
});
