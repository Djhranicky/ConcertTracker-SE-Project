import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';
import { map, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class AuthenticationService {
  private url = 'http://localhost:8080/api';

  private isLoggedIn = false;

  register(email: string, username: string, password: string): Observable<any> {
    const body = {
      email: email,
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
    localStorage.setItem('user', 'paola');
    localStorage.setItem('id', '2');

    return this.http.post(
      `${this.url}/login`,
      { email, password },
      { withCredentials: true }
    );
  }

  logout() {
    // this.isLoggedIn = false;
    document.cookie = 'id=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/api;';
    this.router.navigate(['/login']);
  }

  isAuthenticated(): Observable<boolean> {
    return this.http
      .get<{ message: string }>(`${this.url}/validate`, {
        withCredentials: true,
      })
      .pipe(
        catchError(() => {
          return of(false);
        }),
        map((response) => {
          if (typeof response === 'boolean') {
            return false;
          }
          return response.message === 'user session validated';
        })
      );
  }
  constructor(
    private router: Router,
    private http: HttpClient,
    private cookieService: CookieService
  ) {}
}
