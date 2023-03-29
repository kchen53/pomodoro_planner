import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './components/login/login.component';
import { SessionsComponent } from './pages/sessions/sessions.component';
import { TimerComponent } from './timer/timer.component';
import { TodoListComponent } from './components/todo/todoList.components';

const routes: Routes = [
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
