import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';

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

    return this.http.post(
      `${this.url}/login`,
      { email, password },
      { withCredentials: true }
    );
  }

  logout() {
    // this.isLoggedIn = false;
    localStorage.removeItem('isAuth');
    this.router.navigate(['/login']);
  }

  isAuthenticated(): boolean {
    let isAuth = localStorage.getItem('isAuth');
    return !!isAuth;
  }
  constructor(
    private router: Router,
    private http: HttpClient,
    private cookieService: CookieService
  ) {}
}
