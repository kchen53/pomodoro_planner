import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutUsComponent } from './about-us/about-us.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { TimerComponent } from './timer/timer.component';
import { TodoListComponent } from './todo/todoList.components';

const routes: Routes = [
  { path: 'Pomodoro Planner', component: HomeComponent},
  { path: 'todo', component: TodoListComponent},
  { path: 'about-us', component: AboutUsComponent},
  { path: 'timer', component: TimerComponent},
  { path: 'login', component: LoginComponent},
  { path: 'home', component: HomeComponent},
  { path: '', redirectTo: '/home', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
