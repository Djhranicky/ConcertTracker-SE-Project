import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';

import { PopCarouselComponent } from './pop-carousel.component';

describe('PopCarouselComponent', () => {
  let component: PopCarouselComponent;
  let fixture: ComponentFixture<PopCarouselComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PopCarouselComponent],
      providers: [provideHttpClient(), provideHttpClientTesting()],
    }).compileComponents();

    fixture = TestBed.createComponent(PopCarouselComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
