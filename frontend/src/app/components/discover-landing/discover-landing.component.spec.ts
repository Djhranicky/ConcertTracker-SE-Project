import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiscoverLandingComponent } from './discover-landing.component';

describe('DiscoverLandingComponent', () => {
  let component: DiscoverLandingComponent;
  let fixture: ComponentFixture<DiscoverLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DiscoverLandingComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DiscoverLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
