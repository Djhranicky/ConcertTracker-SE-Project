import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideRouter } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { ConcertComponent } from './concert.component';

describe('ConcertComponent', () => {
  let component: ConcertComponent;
  let fixture: ComponentFixture<ConcertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ConcertComponent],
      providers: [
        provideRouter([]),
        provideHttpClient(),
        provideHttpClientTesting(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(ConcertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
