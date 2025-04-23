// src/app/pages/search/search.component.spec.ts
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Component } from '@angular/core';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { SearchPage } from './search.component';

// Create a mock for the real SearchComponent that doesn't need ActivatedRoute
@Component({
  selector: 'app-search',
  template: '<div class="search-mock">Search Component Mock</div>',
  standalone: true
})
class MockSearchComponent {}

// Create a mock for NavbarComponent
@Component({
  selector: 'app-navbar',
  template: '<div class="navbar-mock">Navbar Mock</div>',
  standalone: true
})
class MockNavbarComponent {}

describe('SearchPage', () => {
  let component: SearchPage;
  let fixture: ComponentFixture<SearchPage>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        SearchPage,
        MockNavbarComponent,
        MockSearchComponent
      ],
      providers: [
        provideHttpClient(),
        provideHttpClientTesting(),
        {
          provide: ActivatedRoute,
          useValue: {
            queryParams: of({ q: 'test' })
          }
        }
      ]
    })
    .compileComponents();

    // Override the component's imports to use our mocks
    TestBed.overrideComponent(SearchPage, {
      set: {
        imports: [MockNavbarComponent, MockSearchComponent]
      }
    });

    fixture = TestBed.createComponent(SearchPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should contain navbar and search components', () => {
    const compiled = fixture.nativeElement;
    
    // Look for our mock components by their class names
    const navbar = compiled.querySelector('.navbar-mock');
    const search = compiled.querySelector('.search-mock');
    
    expect(navbar).not.toBeNull();
    expect(search).not.toBeNull();
  });
});