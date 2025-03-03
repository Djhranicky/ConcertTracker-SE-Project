import { Routes } from '@angular/router';
import { LandingComponent } from './pages/landing/landing.component';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';
import { AuthGuard } from './authentication.guard';
import { GuestGuard } from './guest.guard';

export const appRoutes: Routes = [
  { path: '', component: LandingComponent },
  { path: 'login', component: LoginComponent, canActivate: [GuestGuard] },
  { path: 'register', component: SignupComponent, canActivate: [GuestGuard] },
  { path: '**', component: NotFoundComponent },
];
