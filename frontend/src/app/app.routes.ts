import { Routes } from '@angular/router';
import { LandingComponent } from './pages/landing/landing.component';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';

export const appRoutes: Routes = [
  { path: '', component: LandingComponent },
  //{ path: '**', redirectTo: '/notfound' },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
];
