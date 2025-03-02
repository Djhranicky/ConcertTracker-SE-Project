import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideRouter, RouterModule } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { Location } from '@angular/common';
import { Router } from '@angular/router';
import { SignupComponent } from '../../pages/signup/signup.component';
import { LoginComponent } from '../../pages/login/login.component';
import { LandingHeroComponent } from './landing-hero.component';

describe('LandingHeroComponent', () => {
  let component: LandingHeroComponent;
  let fixture: ComponentFixture<LandingHeroComponent>;
  let router: Router;
  let location: Location;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        LandingHeroComponent,
        RouterModule.forRoot([
          { path: 'register', component: SignupComponent },
          { path: 'login', component: LoginComponent },
        ]),
      ],
      providers: [
        provideRouter([]),
        provideHttpClient(),
        provideHttpClientTesting(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(LandingHeroComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);
    location = TestBed.inject(Location);
    router.initialNavigation();
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should have a button with routerLink set to /register', () => {
    const button = fixture.nativeElement.querySelector('button.signup');
    expect(button.getAttribute('routerLink')).toBe('/register');
  });

  it('should have a button with routerLink set to /login', () => {
    const button = fixture.nativeElement.querySelector('button.login');
    expect(button.getAttribute('routerLink')).toBe('/login');
  });

  it('should navigate to register page when clicking sign up button', async () => {
    const button = fixture.nativeElement.querySelector('button.signup');
    button.click();
    await fixture.whenStable();
    expect(location.path()).toBe('/register');
  });

  it('should navigate to register page when clicking log in button', async () => {
    const button = fixture.nativeElement.querySelector('button.login');
    button.click();
    await fixture.whenStable();
    expect(location.path()).toBe('/login');
  });
});
