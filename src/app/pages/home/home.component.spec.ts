import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HomeComponent } from './home.component';
import { AppModule } from '../../app.module';
import { CalendarComponent } from 'src/app/components/calendar/calendar.component';
import { By } from '@angular/platform-browser';

describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule
      ],
      declarations: [ HomeComponent, CalendarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //New Unit Test

  it('should have an app-calendar component', () => {
    const appCalendar = fixture.debugElement.query(
      By.directive(CalendarComponent)
    );
    expect(appCalendar).not.toBeNull();
  });
  
});
