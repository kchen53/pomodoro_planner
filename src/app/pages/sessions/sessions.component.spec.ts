import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from 'src/app/app.module';
import { SessionsComponent } from './sessions.component';

describe('SessionsComponent', () => {
  let component: SessionsComponent;
  let fixture: ComponentFixture<SessionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule
      ],
      declarations: [ SessionsComponent ]
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
});
