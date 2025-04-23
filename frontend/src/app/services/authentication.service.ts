import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';
import { map, catchError, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class AuthenticationService {
  private url = 'http://localhost:8080/api';

  private isLoggedIn = false;

  register(email: string, username: string, password: string): Observable<any> {
    const body = {
      email: email,
      username: username,
      name: username,
      password: password,
    };
    return this.http.post(`${this.url}/register`, body);
  }

  login(email: string, password: string): Observable<any> {
    const body = {
      email: email,
      password: password,
    };

    localStorage.setItem('isAuth', '1');

    return this.http
      .post<{ username: string }>(
        `${this.url}/login`,
        { email, password },
        { withCredentials: true }
      )
      .pipe(
        tap((response) => {
          if (response && response.username) {
            localStorage.setItem('user', response.username);
          }
        })
      );
  }

  logout() {
    // this.isLoggedIn = false;
    document.cookie = 'id=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/api;';
    localStorage.clear();
    this.router.navigate(['/login']);
  }

  isAuthenticated(): Observable<boolean> {
    let isAuth = localStorage.getItem('isAuth');

    if (isAuth == '1') return of(true);
    return of(false);
  }
  constructor(
    private router: Router,
    private http: HttpClient,
    private cookieService: CookieService
  ) {}
}
