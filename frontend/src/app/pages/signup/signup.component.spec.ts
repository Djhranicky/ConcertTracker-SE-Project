import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideRouter, Router } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { By } from '@angular/platform-browser';
import { of, throwError } from 'rxjs';

import { SignupComponent } from './signup.component';
import { AuthenticationService } from '../../services/authentication.service';

describe('SignupComponent', () => {
  let component: SignupComponent;
  let fixture: ComponentFixture<SignupComponent>;
  let service: AuthenticationService;
  let authMock: jasmine.SpyObj<AuthenticationService>;
  let routerMock: jasmine.SpyObj<Router>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SignupComponent, ReactiveFormsModule],
      providers: [
        provideRouter([]),
        provideHttpClient(),
        provideHttpClientTesting(),
        AuthenticationService,
      ],
      teardown: { destroyAfterEach: false },
    }).compileComponents();

    service = TestBed.inject(AuthenticationService);
    authMock = jasmine.createSpyObj('AuthenticationService', ['login']);
    routerMock = jasmine.createSpyObj('Router', ['navigate']);
    fixture = TestBed.createComponent(SignupComponent);
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

    const userInput = fixture.debugElement.query(By.css('input[type="text"]'));

    const passwordInput = fixture.debugElement.query(
      By.css('input[type="password"]')
    );
    const loginButton = fixture.debugElement.query(
      By.css('button[type="submit"]')
    );

    expect(usernameInput).toBeTruthy();
    expect(userInput).toBeTruthy();
    expect(passwordInput).toBeTruthy();
    expect(loginButton).toBeTruthy();
  });

  it('should have email, user and passwords as required fields of form', () => {
    expect(component.signupForm).toBeDefined();
    expect(component.signupForm.contains('email')).toBeTrue();
    expect(component.signupForm.contains('username')).toBeTrue();
    expect(component.signupForm.contains('password')).toBeTrue();
    expect(component.signupForm.contains('confirmPassword')).toBeTrue();
  });

  //test email validators
  it('should require email field', () => {
    const email = component.signupForm.get('email');
    email?.setValue('');
    expect(email?.valid).toBeFalse();
    expect(email?.errors?.['required']).toBeTruthy();
  });

  it('should validate email field to be email format', () => {
    const email = component.signupForm.get('email');
    email?.setValue('hjdfsb');
    expect(email?.valid).toBeFalse();
    expect(email?.errors?.['email']).toBeTruthy();

    email?.setValue('email@email.com');
    expect(email?.valid).toBeTrue();
  });

  //test password validators
  it('should require password field', () => {
    const password = component.signupForm.get('password');
    password?.setValue('');
    expect(password?.valid).toBeFalse();
    expect(password?.errors?.['required']).toBeTruthy();
  });

  //test confirm passwords

  //test register() call
  it('should not call AuthenticationService.register if form is invalid', () => {
    component.signupForm.setValue({
      email: 'dfhgerg',
      username: '23234',
      password: '',
      confirmPassword: '',
    });

    component.register();
    expect(authMock.login).not.toHaveBeenCalled();
  });
});
