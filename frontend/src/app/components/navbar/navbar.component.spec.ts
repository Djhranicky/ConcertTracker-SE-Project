import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideRouter, RouterModule } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { Location } from '@angular/common';
import { Router } from '@angular/router';
import { SignupComponent } from '../../pages/signup/signup.component';
import { LoginComponent } from '../../pages/login/login.component';
import { NavbarComponent } from './navbar.component';
import { AuthenticationService } from '../../services/authentication.service';
import { of } from 'rxjs';

describe('NavbarComponent', () => {
  let component: NavbarComponent;
  let fixture: ComponentFixture<NavbarComponent>;
  let router: Router;
  let location: Location;
  let authenticationService: jasmine.SpyObj<AuthenticationService>;

  beforeEach(async () => {
    const authServiceSpy = jasmine.createSpyObj('AuthenticationService', [
      'isAuthenticated',
      'logout',
    ]);

    authServiceSpy.isAuthenticated.and.returnValue(of(false));

    await TestBed.configureTestingModule({
      imports: [
        NavbarComponent,
        RouterModule.forRoot([
          { path: 'register', component: SignupComponent },
          { path: 'login', component: LoginComponent },
        ]),
      ],
      providers: [
        provideRouter([]),
        provideHttpClient(),
        provideHttpClientTesting(),
        { provide: AuthenticationService, useValue: authServiceSpy },
      ],
      // teardown: { destroyAfterEach: false },
    }).compileComponents();

    router = TestBed.inject(Router);
    location = TestBed.inject(Location);
    router.initialNavigation();
    authenticationService = TestBed.inject(
      AuthenticationService
    ) as jasmine.SpyObj<AuthenticationService>;
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should have a button with routerLink set to /register when logged out', () => {
    const button = fixture.nativeElement.querySelector('p-button.register');
    expect(button.getAttribute('routerLink')).toBe('/register');
  });

  it('should have a button with routerLink set to /login  when logged out', () => {
    const button = fixture.nativeElement.querySelector('p-button.log-in');
    expect(button.getAttribute('routerLink')).toBe('/login');
  });

  it('should navigate to register page when clicking sign up button  when logged out', async () => {
    const button = fixture.nativeElement.querySelector('p-button.register');
    button.click();
    await fixture.whenStable();
    expect(location.path()).toBe('/register');
  });

  it('should navigate to register page when clicking log in button  when logged out', async () => {
    const button = fixture.nativeElement.querySelector('p-button.log-in');
    button.click();
    await fixture.whenStable();
    expect(location.path()).toBe('/login');
  });

  it('should log out if logged in', () => {
    authenticationService.logout.and.stub();
    component.logout();
    expect(authenticationService.logout).toHaveBeenCalled();
  });

  it('should show menubar items if logged in', () => {
    authenticationService.isAuthenticated.and.returnValue(of(true));
    component.ngOnInit();
    expect(component.isLoggedIn).toBe(true);
    expect(component.items?.length).toBeGreaterThan(0);
  });

  it('should not show menubar items if not logged in', () => {
    component.logout();
    authenticationService.isAuthenticated.and.returnValue(of(false));
    component.ngOnInit();
    expect(component.isLoggedIn).toBeFalse();
    expect(component.items).toBeUndefined();
  });
});
