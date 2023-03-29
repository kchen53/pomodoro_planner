import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatToolbarModule} from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { NavbarComponent } from './components/navbar/navbar.component';
import { LoginComponent } from './components/login/login.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { HttpClientModule } from '@angular/common/http';
import { HomeComponent } from './pages/home/home.component';
import { TodoListComponent } from './components/todo/todoList.components';
import { MatListModule } from '@angular/material/list';
import { MatIconModule} from '@angular/material/icon';
import { TodoService } from './components/todo/todoService';
import { TimerComponent } from './timer/timer.component';
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { MatDialogModule, MAT_DIALOG_DATA } from '@angular/material/dialog';
import {MatCheckboxModule} from '@angular/material/checkbox';
import { SessionsComponent } from './pages/sessions/sessions.component';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { DailyTaskComponent } from './components/daily-task/daily-task.component';
import { CalendarComponent } from './components/calendar/calendar.component';
import { MatNativeDateModule } from '@angular/material/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { CalendarModule, DateAdapter } from 'angular-calendar';
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns';
import { SpotifyApiComponent } from './components/spotify-api/spotify-api.component';
import {MatDialogRef} from '@angular/material/dialog';


@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    LoginComponent,
    HomeComponent,
    TodoListComponent,
    TimerComponent,
    SessionsComponent,
    DailyTaskComponent,
    CalendarComponent,
    SpotifyApiComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatButtonModule,
    MatCardModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatListModule,
    MatIconModule,
    FormsModule,
    MatDialogModule,
    MatCheckboxModule,
    MatProgressBarModule,
    MatNativeDateModule,
    MatDatepickerModule,
    CalendarModule.forRoot({
      provide: DateAdapter,
      useFactory: adapterFactory,
    }),
    MatDialogModule
  ],

  providers: [TodoService],
  bootstrap: [AppComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]

})
export class AppModule { }

