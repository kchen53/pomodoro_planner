import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { SessionsComponent } from './sessions/sessions.component';
import { TimerComponent } from '../timer/timer.component';
import { TodoListComponent } from './todo/todoList.components';

const routes: Routes = [
  { path: 'Pomodoro Planner', component: HomeComponent},
  { path: 'todo', component: TodoListComponent},
  { path: 'timer', component: TimerComponent},
  { path: 'login', component: LoginComponent},
  { path: 'home', component: HomeComponent},
  { path: 'session', component: SessionsComponent},
  { path: '', redirectTo: '/home', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
