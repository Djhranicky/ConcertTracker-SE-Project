import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PopCarouselComponent } from './pop-carousel.component';

describe('PopCarouselComponent', () => {
  let component: PopCarouselComponent;
  let fixture: ComponentFixture<PopCarouselComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PopCarouselComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PopCarouselComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
