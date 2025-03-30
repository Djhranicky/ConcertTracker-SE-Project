import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { UserService } from './user.service';

describe('UserService', () => {
  let service: UserService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting()]
    });
    service = TestBed.inject(UserService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should return user profile', () => {
    service.getUserProfile().subscribe(profile => {
      expect(profile).toBeTruthy();
      expect(profile.name).toBeDefined();
      expect(profile.username).toBeDefined();
      expect(profile.bio).toBeDefined();
      expect(profile.profileImage).toBeDefined();
      expect(profile.stats).toBeDefined();
    });
  });

  it('should return favorite concerts', () => {
    service.getFavoriteConcerts().subscribe(concerts => {
      expect(concerts.length).toBeGreaterThan(0);
      concerts.forEach(concert => {
        expect(concert.title).toBeDefined();
        expect(concert.artist).toBeDefined();
        expect(concert.date).toBeDefined();
        expect(concert.image).toBeDefined();
      });
    });
  });

  it('should return user posts', () => {
    service.getUserPosts().subscribe(posts => {
      expect(posts.length).toBeGreaterThan(0);
      posts.forEach(post => {
        expect(post.type).toBeDefined();
        expect(post.username).toBeDefined();
        expect(post.avatar).toBeDefined();
        expect(post.artist).toBeDefined();
        expect(post.tour).toBeDefined();
        expect(post.img).toBeDefined();
      });
    });
  });

  it('should convert concert card to post format', () => {
    const concertCard = {
      title: 'Test Tour',
      artist: 'Test Artist',
      date: 'Jan 1, 2025',
      image: 'test-image.jpg'
    };
    
    const post = service.convertToPost(concertCard, 'review');
    
    expect(post.type).toBe('review');
    expect(post.tour).toBe(concertCard.title);
    expect(post.artist).toBe(concertCard.artist);
    expect(post.date).toBe(concertCard.date);
    expect(post.img).toBe(concertCard.image);
  });
});