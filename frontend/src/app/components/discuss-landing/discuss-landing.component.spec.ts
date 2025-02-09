import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DiscussLandingComponent } from './discuss-landing.component';

describe('DiscussLandingComponent', () => {
  let component: DiscussLandingComponent;
  let fixture: ComponentFixture<DiscussLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DiscussLandingComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DiscussLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
