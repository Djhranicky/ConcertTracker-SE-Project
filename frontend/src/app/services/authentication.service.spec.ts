import { TestBed } from '@angular/core/testing';
import {
  provideHttpClientTesting,
  HttpTestingController,
} from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';

import { AuthenticationService } from './authentication.service';

describe('AuthenticationService', () => {
  let service: AuthenticationService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        AuthenticationService,
        provideHttpClient(),
        provideHttpClientTesting(),
      ],
      teardown: { destroyAfterEach: false },
    });
    service = TestBed.inject(AuthenticationService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  //register test success
  it('should send a POST to /register endpoint in backend', (done) => {
    const mockRegister = {
      email: 'mock@test.com',
      password: 'password123',
      name: 'mocktest',
    };

    const mockResponse = {};
    service
      .register(mockRegister.email, mockRegister.name, mockRegister.password)
      .subscribe((response) => {
        expect(response).toEqual(mockResponse);
        done();
      });

    const req = httpMock.expectOne(`${service['url']}/register`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(mockRegister);

    req.flush(mockResponse, { status: 200, statusText: 'OK' });
  });

  //test register error (user already exists)
  it('should handle existing user register error', () => {
    const mockRegister = {
      email: 'test@test.com',
      password: 'password123',
      name: 'mocktest',
    };

    service
      .register(mockRegister.email, mockRegister.name, mockRegister.password)
      .subscribe({
        next: () => {
          fail('Expected registration to fail');
        },
        error: (error: HttpErrorResponse) => {
          expect(error.status).toBe(400);
        },
      });

    const req = httpMock.expectOne(`${service['url']}/register`);
    req.flush(null, { status: 400, statusText: '' });
  });

  //login test
  it('should send a POST to /login endpoint in backend', (done) => {
    const mockLogin = {
      email: 'mock@test.com',
      password: 'password123',
    };

    const mockResponse = {};
    service.login(mockLogin.email, mockLogin.password).subscribe((response) => {
      expect(response).toEqual(mockResponse);
      done();
    });

    const req = httpMock.expectOne(`${service['url']}/login`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(mockLogin);

    req.flush(mockResponse, { status: 200, statusText: 'OK' });
  });

  //test login error (wrong password or user does not exist)
  it('should handle existing user login error', () => {
    const mockLogin = {
      email: 'test@test.com',
      password: 'password',
    };

    service.login(mockLogin.email, mockLogin.password).subscribe({
      next: () => {
        fail('Expected login to fail');
      },
      error: (error: HttpErrorResponse) => {
        expect(error.status).toBe(400);
      },
    });

    const req = httpMock.expectOne(`${service['url']}/login`);
    req.flush(null, { status: 400, statusText: '' });
  });

  //logout test
  it('should delete session from localStorage', () => {
    localStorage.setItem('isAuth', '1');
    service.logout();
    expect(localStorage.getItem('isAuth')).toBeNull();
  });

  //isAuth test
  it('should return true if session exists in localStorage', () => {
    localStorage.setItem('isAuth', '1');
    expect(service.isAuthenticated()).toBe(true);
  });

  it('should return false if session does not exist in localStorage', () => {
    service.logout();
    expect(service.isAuthenticated()).toBe(false);
  });
});
