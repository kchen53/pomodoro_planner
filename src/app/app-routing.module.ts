import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutUsComponent } from './about-us/about-us.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { TimerComponent } from './timer/timer.component';

const routes: Routes = [
  { path: 'Pomodoro Planner', component: HomeComponent},
  { path: 'home', component: HomeComponent},
  { path: 'about-us', component: AboutUsComponent},
  { path: 'timer', component: TimerComponent},
  { path: 'login', component: LoginComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
