import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SpotifyApiComponent } from './spotify-api.component';
import { AppModule } from '../../app.module';

describe('SpotifyApiComponent', () => {
  let component: SpotifyApiComponent;
  let fixture: ComponentFixture<SpotifyApiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule
      ],
      declarations: [ SpotifyApiComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SpotifyApiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //New Unit Test
});
