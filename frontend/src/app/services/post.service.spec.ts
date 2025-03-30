import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { PostService, Post } from './post.service';

describe('PostService', () => {
  let service: PostService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting()],
    });
    service = TestBed.inject(PostService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should return an observable of posts', () => {
    const mockPosts: Post[] = [
      {
        type: 'review',
        username: 'John Doe',
        avatar: 'imgurl',
        postDate: 'Feb 20, 2025',
        date: 'Feb 19, 2025',
        venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
        artist: 'Billie Eilish',
        tour: 'HIT ME HARD AND SOFT',
        img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
        reviewText: null,
        attachedImg: null,
        setlist: null,
        rating: 4,
        likes: 10,
        comments: 3,
      },
      {
        type: 'wishlist',
        username: 'Jane Doe',
        avatar: 'imgurl',
        postDate: 'Feb 20, 2025',
        artist: 'Bad Bunny',
        tour: 'No me quiero ir de aqui',
        img: 'https://i.scdn.co/image/ab6761610000e5eb81f47f44084e0a09b5f0fa13',
        date: null,
        venue: null,
        reviewText: null,
        attachedImg: null,
        rating: 4,
        setlist: null,
        likes: 2,
        comments: 0,
      },
      {
        type: 'review',
        username: 'Jane Doe',
        avatar: 'imgurl',
        postDate: 'Feb 20, 2025',
        date: 'Feb 19, 2025',
        venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
        artist: 'Billie Eilish',
        tour: 'HIT ME HARD AND SOFT',
        img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
        reviewText:
          'Lorem ipsum odor amet, consectetuer adipiscing elit. Potenti mus fermentum sed dapibus egestas; aptent faucibus quisque? Fames fringilla consectetur tortor leo potenti at porttitor aenean. Vehicula sociosqu nam in litora malesuada. Lacinia quisque gravida imperdiet magnis magna lacinia senectus. Vestibulum morbi netus nullam; parturient nostra tellus posuere non.',
        attachedImg: null,
        setlist: null,
        rating: 5,
        likes: 2,
        comments: 0,
      },
    ];

    service.getPosts().subscribe((posts) => {
      expect(posts).toEqual(mockPosts);
    });
  });

  it('should return an observable of the correct type', () => {
    service.getPosts().subscribe((posts) => {
      expect(Array.isArray(posts)).toBe(true);
      posts.forEach((post) => {
        expect(post.type).toBeDefined();
        expect(post.username).toBeDefined();
        expect(post.avatar).toBeDefined();
        expect(post.postDate).toBeDefined();
        expect(post.date).toBeDefined();
        expect(post.venue).toBeDefined();
        expect(post.artist).toBeDefined();
        expect(post.tour).toBeDefined();
        expect(post.img).toBeDefined();
        expect(post.reviewText).toBeDefined();
        expect(post.attachedImg).toBeDefined();
        expect(post.rating).toBeDefined();
        expect(post.likes).toBeDefined();
        expect(post.comments).toBeDefined();
        expect(post.setlist).toBeDefined();
      });
    });
  });
});
