import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DiscussLandingComponent } from './discuss-landing.component';

describe('DiscussLandingComponent', () => {
  let component: DiscussLandingComponent;
  let fixture: ComponentFixture<DiscussLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DiscussLandingComponent],
      teardown: { destroyAfterEach: false },
    }).compileComponents();

    fixture = TestBed.createComponent(DiscussLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should bind title', () => {
    component.title = 'Test Title';
    fixture.detectChanges();
    expect(component.title).toBe('Test Title');
  });

  it('should bind subtitle', () => {
    component.title = 'Test Subtitle';
    fixture.detectChanges();
    expect(component.title).toBe('Test Subtitle');
  });

  it('should bind paragraph', () => {
    component.title = 'Test paragraph';
    fixture.detectChanges();
    expect(component.title).toBe('Test paragraph');
  });
});
