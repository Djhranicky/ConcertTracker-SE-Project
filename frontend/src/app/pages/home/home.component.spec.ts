import { ComponentFixture, fakeAsync, TestBed } from '@angular/core/testing';
import { provideRouter } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { HomeComponent } from './home.component';
import { of } from 'rxjs';
import { AuthenticationService } from '../../services/authentication.service';

describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;
  let authenticationService: jasmine.SpyObj<AuthenticationService>;
  beforeEach(async () => {
    authenticationService = jasmine.createSpyObj('AuthenticationService', [
      'isAuthenticated',
    ]);

    await TestBed.configureTestingModule({
      imports: [HomeComponent],
      providers: [
        provideRouter([]),
        provideHttpClient(),
        provideHttpClientTesting(),
        { provide: AuthenticationService, useValue: authenticationService },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
    // fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should set isLoggedIn true when auth valid', () => {
    authenticationService.isAuthenticated.and.returnValue(of(true));
    fixture.detectChanges();
    expect(component.isLoggedIn).toBeTrue();
    // expect(authenticationService.isAuthenticated).toHaveBeenCalled();
  });

  it('should set isLoggedIn false when auth invalid', () => {
    authenticationService.isAuthenticated.and.returnValue(of(false));
    fixture.detectChanges();
    expect(component.isLoggedIn).toBeFalse();
  });
});
