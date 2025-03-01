import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class AuthenticationService {
  private url = 'http://localhost:8080/api';

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
    return this.http.post(`${this.url}/login`, { email, password });
  }

  storeJWT(jwt: string): void {
    localStorage.setItem('jwt', jwt);
  }

  getJWT(): string | null {
    return localStorage.getItem('jwt');
  }

  logout() {
    localStorage.removeItem('jwt');
    this.router.navigate(['/login']);
  }

  isAuthenticated(): boolean {
    return !!this.getJWT();
  }
  constructor(private router: Router, private http: HttpClient) {}
}
