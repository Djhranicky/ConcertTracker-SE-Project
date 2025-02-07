import { Routes } from '@angular/router';
import { LandingComponent } from './pages/landing/landing.component';

export const appRoutes: Routes = [
  { path: '', component: LandingComponent },
  //{ path: '**', redirectTo: '/notfound' },
];
