import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from 'src/app/app.module';
import { SessionsComponent } from './sessions.component';
import { DailyTaskComponent } from 'src/app/components/daily-task/daily-task.component';
import { By } from '@angular/platform-browser';

describe('SessionsComponent', () => {
  let component: SessionsComponent;
  let fixture: ComponentFixture<SessionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule
      ],
      declarations: [ SessionsComponent, DailyTaskComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SessionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //New Unit Test
  it('should have an app-dailyTask component', () => {
    const appCalendar = fixture.debugElement.query(
      By.directive(DailyTaskComponent)
    );
    expect(appCalendar).not.toBeNull();
  });
});
