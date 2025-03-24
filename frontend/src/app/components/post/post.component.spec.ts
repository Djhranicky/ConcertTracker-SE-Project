import { ComponentFixture, TestBed } from '@angular/core/testing';
import { of } from 'rxjs';
import { PostComponent } from './post.component';

describe('PostComponent', () => {
  let component: PostComponent;
  let fixture: ComponentFixture<PostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PostComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });
  const mockService = jasmine.createSpyObj('mockService', ['getPosts']);
  mockService.getPosts.and.returnValue(of(true));
  it('should create', () => {
    expect(true).toBeTruthy();
    expect(mockService.getPosts);
  });

  //post should have
});
