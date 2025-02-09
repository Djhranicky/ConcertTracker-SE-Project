import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FeaturesLandingComponent } from './features-landing.component';

describe('FeaturesLandingComponent', () => {
  let component: FeaturesLandingComponent;
  let fixture: ComponentFixture<FeaturesLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FeaturesLandingComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(FeaturesLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
