import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiscoverLandingComponent } from './discover-landing.component';

describe('DiscoverLandingComponent', () => {
  let component: DiscoverLandingComponent;
  let fixture: ComponentFixture<DiscoverLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DiscoverLandingComponent],
      teardown: { destroyAfterEach: false },
    }).compileComponents();

    fixture = TestBed.createComponent(DiscoverLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
