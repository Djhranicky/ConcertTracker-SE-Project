import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { HttpTestingController } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';

import { SearchComponent } from './search.component';

describe('SearchComponent', () => {
  let component: SearchComponent;
  let fixture: ComponentFixture<SearchComponent>;
  let httpMock: HttpTestingController;
  let activatedRouteMock: any;

  beforeEach(async () => {
    // Create a mock for ActivatedRoute
    activatedRouteMock = {
      queryParams: of({ q: 'testArtist' })
    };

    await TestBed.configureTestingModule({
      imports: [
        SearchComponent,
        RouterTestingModule
      ],
      providers: [
        provideHttpClient(),
        provideHttpClientTesting(),
        { provide: ActivatedRoute, useValue: activatedRouteMock }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(SearchComponent);
    component = fixture.componentInstance;
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    // Verify there are no outstanding HTTP requests
    httpMock.verify();
  });

  it('should create', () => {
    // Handle the HTTP request triggered by ngOnInit
    fixture.detectChanges();
    const req = httpMock.expectOne('http://localhost:8080/api/artist?name=testArtist');
    req.flush({ artist: { name: 'Test Artist' } });
    
    expect(component).toBeTruthy();
  });

  it('should get query parameter and search for artist on init', () => {
    const searchSpy = spyOn(component, 'searchArtists').and.callThrough();
    
    fixture.detectChanges(); 
    
    expect(component.query).toBe('testArtist');
    expect(searchSpy).toHaveBeenCalled();
    const req = httpMock.expectOne('http://localhost:8080/api/artist?name=testArtist');
    req.flush({ artist: { name: 'Test Artist' } });
  });


  it('should show error message when error exists and not loading', () => {
    fixture.detectChanges();
    const req = httpMock.expectOne('http://localhost:8080/api/artist?name=testArtist');
    req.flush({ artist: { name: 'Test Artist' } });
    component.error = 'Test error';
    component.loading = false;
    fixture.detectChanges();
    
    const errorMsg = fixture.nativeElement.querySelector('.error-message');
    expect(errorMsg).toBeTruthy();
    expect(errorMsg.textContent).toContain('Test error');
  });

  it('should display search results when data is loaded', () => {
    fixture.detectChanges();
    const req = httpMock.expectOne('http://localhost:8080/api/artist?name=testArtist');
    req.flush({ artist: { name: 'Test Artist' } });
    component.loading = false;
    component.results = [{ artist: { name: 'Test Artist' } }];
    fixture.detectChanges();
  
    const card = fixture.nativeElement.querySelector('p-card');
    expect(card).toBeTruthy();
  });

  it('should make HTTP request with correct URL when searching', fakeAsync(() => {
    component.query = 'searchTestArtist';
    component.searchArtists();
    const req = httpMock.expectOne('http://localhost:8080/api/artist?name=searchTestArtist');
    expect(req.request.method).toBe('GET');
    
    // Respond with mock data
    req.flush({
      artist: {
        name: 'Search Test Artist'
      }
    });
    
    // Expectations
    expect(component.loading).toBeFalse();
    expect(component.results.length).toBe(1);
  }));

  it('should set error message when HTTP request fails', fakeAsync(() => {
    component.query = 'nonExistentArtist';
    component.searchArtists();
    const req = httpMock.expectOne('http://localhost:8080/api/artist?name=nonExistentArtist');
    req.error(new ErrorEvent('Network error'));
    expect(component.loading).toBeFalse();
    expect(component.error).toBe('No results found.');
  }));

  it('should initialize results array when there is no query', fakeAsync(() => {
    TestBed.resetTestingModule();
    
    const emptyQueryMock = {
      queryParams: of({})
    };
    
    TestBed.configureTestingModule({
      imports: [
        SearchComponent,
        RouterTestingModule
      ],
      providers: [
        provideHttpClient(),
        provideHttpClientTesting(),
        { provide: ActivatedRoute, useValue: emptyQueryMock }
      ]
    }).compileComponents();
    
    const newFixture = TestBed.createComponent(SearchComponent);
    const newComponent = newFixture.componentInstance;
    
    // Spy on the searchArtists method
    const searchSpy = spyOn(newComponent, 'searchArtists');
    
    newFixture.detectChanges(); // This triggers ngOnInit
    
    expect(newComponent.query).toBe('');
    expect(searchSpy).not.toHaveBeenCalled();
    expect(newComponent.results).toEqual([]);
  }));
});
