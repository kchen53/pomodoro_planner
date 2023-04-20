import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from 'src/app/app.module';
import { DailyTaskComponent } from './daily-task.component';

describe('DailyTaskComponent', () => {
  let component: DailyTaskComponent;
  let fixture: ComponentFixture<DailyTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule
      ],
      declarations: [ DailyTaskComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DailyTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //New Unit Test
  it('check if the daily calender displays the correct date', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('h2').textContent).toContain('Wednesday, April 19, 2023');
  });
});
