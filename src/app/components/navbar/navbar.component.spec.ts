import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { By } from '@angular/platform-browser';
import { RouterTestingModule } from '@angular/router/testing';
import { AppModule } from '../../app.module';
import { NavbarComponent } from './navbar.component';

describe('NavbarComponent', () => {
  let component: NavbarComponent;
  let fixture: ComponentFixture<NavbarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule, 
        MatButtonModule, 
        RouterTestingModule,
        MatIconModule
      ],
      declarations: [ NavbarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //New Unit Test

  it('should have a link to the session page', () => {
    const link = fixture.debugElement.query(By.css('a[routerLink="session"]'));
    expect(link).toBeTruthy();
  });

  it('should have a settings icon', () => {
    const icon = fixture.debugElement.query(By.css('mat-icon'));
    expect(icon.nativeElement.innerText).toContain('settings');
  });

});
