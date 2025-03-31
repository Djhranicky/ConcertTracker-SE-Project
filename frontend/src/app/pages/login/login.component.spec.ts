import {
  ComponentFixture,
  fakeAsync,
  TestBed,
  tick,
} from '@angular/core/testing';
import { provideRouter, Router } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { By } from '@angular/platform-browser';
import { of, throwError } from 'rxjs';

import { LoginComponent } from './login.component';
import { AuthenticationService } from '../../services/authentication.service';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let service: AuthenticationService;
  let authMock: jasmine.SpyObj<AuthenticationService>;
  let routerMock: jasmine.SpyObj<Router>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoginComponent, ReactiveFormsModule],
      providers: [
        provideRouter([]),
        provideHttpClient(),
        provideHttpClientTesting(),
        AuthenticationService,
      ],
    }).compileComponents();

    service = TestBed.inject(
      AuthenticationService
    ) as jasmine.SpyObj<AuthenticationService>;
    authMock = jasmine.createSpyObj('AuthenticationService', ['login']);
    routerMock = jasmine.createSpyObj('Router', ['navigate']);
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should render username and password input fields and login button', () => {
    const usernameInput = fixture.debugElement.query(
      By.css('input[type="email"]')
    );
    const passwordInput = fixture.debugElement.query(
      By.css('input[type="password"]')
    );
    const loginButton = fixture.debugElement.query(
      By.css('button[type="submit"]')
    );

    expect(usernameInput).toBeTruthy();
    expect(passwordInput).toBeTruthy();
    expect(loginButton).toBeTruthy();
  });

  it('should have email and passwords as required fields of form', () => {
    expect(component.loginForm).toBeDefined();
    expect(component.loginForm.contains('email')).toBeTrue();
    expect(component.loginForm.contains('password')).toBeTrue();
  });

  //test email validators
  it('should require email field', () => {
    const email = component.loginForm.get('email');
    email?.setValue('');
    expect(email?.valid).toBeFalse();
    expect(email?.errors?.['required']).toBeTruthy();
  });

  it('should validate email field to be email format', () => {
    const email = component.loginForm.get('email');
    email?.setValue('hjdfsb');
    expect(email?.valid).toBeFalse();
    expect(email?.errors?.['email']).toBeTruthy();

    email?.setValue('email@email.com');
    expect(email?.valid).toBeTrue();
  });

  //test password validators
  it('should require password field', () => {
    const password = component.loginForm.get('password');
    password?.setValue('');
    expect(password?.valid).toBeFalse();
    expect(password?.errors?.['required']).toBeTruthy();
  });

  //test login() call
  it('should call login() if form valid', () => {
    component.loginForm.controls['email'].setValue('test@test.com');
    component.loginForm.controls['password'].setValue('password123');
    expect(component.loginForm.value).toEqual({
      email: 'test@test.com',
      password: 'password123',
    });

    //test click button
    //test /dashboard redirect
  });

  it('should not call AuthenticationService.login if form is invalid', () => {
    component.loginForm.setValue({
      email: 'dfhgerg',
      password: '',
    });

    component.login();
    expect(authMock.login).not.toHaveBeenCalled();
  });

  it('should call AuthenticationService.login if form is valid', () => {
    spyOn(service, 'login').and.returnValue(of({}));
    component.loginForm.setValue({
      email: 'test@test.com',
      password: 'password123',
    });

    component.login();
    expect(service.login).toHaveBeenCalledWith('test@test.com', 'password123');
  });

  it('should show alert if login fails', fakeAsync(() => {
    const error = { message: 'Error message' };
    spyOn(service, 'login').and.returnValue(throwError(() => error));
    spyOn(window, 'alert');

    component.ngOnInit();
    component.loginForm.setValue({
      email: 'test@test.com',
      password: 'password123',
    });

    component.login();
    tick();
    expect(window.alert).toHaveBeenCalledWith('Login failed: ' + error.message);
  }));
});
