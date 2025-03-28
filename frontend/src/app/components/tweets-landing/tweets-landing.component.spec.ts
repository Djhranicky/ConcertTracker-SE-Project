import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TweetsLandingComponent } from './tweets-landing.component';

describe('TweetsLandingComponent', () => {
  let component: TweetsLandingComponent;
  let fixture: ComponentFixture<TweetsLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TweetsLandingComponent],
      teardown: { destroyAfterEach: false },
    }).compileComponents();

    fixture = TestBed.createComponent(TweetsLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
