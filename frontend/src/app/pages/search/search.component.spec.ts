import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SearchPage } from './search.component'; // Changed from SearchComponent to SearchPage

describe('SearchComponent', () => {
  let component: SearchPage; 
  let fixture: ComponentFixture<SearchPage>; 
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SearchPage] 
    })
    .compileComponents();

    fixture = TestBed.createComponent(SearchPage); 
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});