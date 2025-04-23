import {
  ComponentFixture,
  TestBed,
  fakeAsync,
  tick,
} from '@angular/core/testing';
import { of } from 'rxjs';
import { PostComponent } from './post.component';
import { Post } from '../../models/post.model';

import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
describe('PostComponent', () => {
  let component: PostComponent;
  let fixture: ComponentFixture<PostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PostComponent],
      providers: [provideHttpClient(), provideHttpClientTesting()],
      teardown: { destroyAfterEach: false },
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PostComponent);
    component = fixture.componentInstance;
    component.post = {
      type: 'review',
      username: 'John Doe',
      avatar: 'imgurl',
      postDate: 'Feb 20, 2025',
      date: 'Feb 19, 2025',
      venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
      artist: 'Billie Eilish',
      setlist: null,
      tour: 'HIT ME HARD AND SOFT',
      img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
      reviewText: null,
      attachedImg: null,
      rating: 4,
      likes: 10,
      comments: 3,
      city: 'Brisbane, Australia', // Added missing property
      id: '3'
    };
    fixture.detectChanges();
  });
  // const mockService = jasmine.createSpyObj('mockService', ['getPosts']);
  // mockService.getPosts.and.returnValue(of(true));
  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //toggleLike()
  it('should toggle isLiked affect post like count', () => {
    component.toggleLike();
    expect(component.isLiked).toBeTrue();
    expect(component.post.likes).toBe(11);

    component.toggleLike();
    expect(component.isLiked).toBeFalse();
    expect(component.post.likes).toBe(10);
  });
});
