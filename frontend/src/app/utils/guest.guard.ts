import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { AuthenticationService } from '../services/authentication.service';
import { catchError, Observable, map } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class GuestGuard implements CanActivate {
  constructor(
    private authService: AuthenticationService,
    private router: Router
  ) {}

  canActivate(): Observable<boolean> {
    // if (!this.authService.isAuthenticated()) {
    //   return true;
    // } else {
    //   this.router.navigate(['/']);
    //   return false;
    // }
    return this.authService.isAuthenticated().pipe(
      map((isAuth) => {
        if (isAuth) {
          this.router.navigate(['/']);
          return false;
        } else {
          return true;
        }
      }),
      catchError(() => {
        return [true];
      })
    );
  }
}
