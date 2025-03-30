import { Routes } from '@angular/router';
import { LandingComponent } from './pages/landing/landing.component';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { UserProfileComponent } from './pages/user-profile/user-profile.component';
import { AuthGuard } from './utils/authentication.guard';
import { GuestGuard } from './utils/guest.guard';
import { HomeComponent } from './pages/home/home.component';
import { ConcertComponent } from './pages/concert/concert.component';
export const appRoutes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'login', component: LoginComponent, canActivate: [GuestGuard] },
  { path: 'register', component: SignupComponent, canActivate: [GuestGuard] },
  { path: 'concerts', component: ConcertComponent },
  {
    path: 'user-profile',
    component: UserProfileComponent,
    canActivate: [AuthGuard],
  },
  { path: '**', component: NotFoundComponent },
];
