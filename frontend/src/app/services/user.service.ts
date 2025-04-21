import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Post } from './post.service';
import { Concert, Tour } from './concert.service';

export interface UserProfile {
  name: string;
  username: string;
  bio: string;
  profileImage: string;
  stats: {
    concerts: number;
    lists: number;
    following: number;
    followers: number;
  };
}

export interface ConcertCard {
  title: string;
  artist: string;
  date: string;
  image: string;
}

export interface Activity {
  text: string;
}

export interface List {
  title: string;
  thumbnails: string[];
}

@Injectable({
  providedIn: 'root',
})
export class UserService {
  // Mock user profile data
  userProfile: UserProfile = {
    name: 'Jane Smith',
    username: 'janesmith',
    bio: '24. music lover. love pop music.',
    profileImage: 'imgs/user-profile.jpeg',
    stats: {
      concerts: 23,
      lists: 3,
      following: 21,
      followers: 19,
    },
  };

  // Mock data for favorite concerts
  favoriteConcerts: ConcertCard[] = [
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
  ];

  // Mock data for recent attendance
  recentAttendance: ConcertCard[] = [
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
  ];

  // Mock data for bucket list
  bucketList: ConcertCard[] = [
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    },
  ];

  // Mock data for recent activity
  recentActivity: Activity[] = [
    {
      text: 'You followed <span class="highlight">John Doe</span>',
    },
    {
      text: 'You added a show to <span class="highlight">2025 shows</span> list',
    },
    {
      text: 'You attended Billie Eilish\'s <span class="highlight">HIT ME HARD AND SOFT</span>',
    },
  ];

  // Mock data for recent lists
  recentLists: List[] = [
    {
      title: '2025 Shows',
      thumbnails: [
        'imgs/post-malone.png',
        'imgs/starboy.png',
        'imgs/linkin.png',
      ],
    },
    {
      title: 'Festivals',
      thumbnails: [
        'imgs/post-malone.png',
        'imgs/starboy.png',
        'imgs/linkin.png',
      ],
    },
  ];

  // User posts for timeline
  userPosts: Post[] = [
    {
      type: 'review',
      username: 'Jane Smith',
      avatar: 'imgs/user-profile.jpeg',
      postDate: 'Mar 30, 2025',
      date: 'Feb 19, 2025',
      venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
      artist: 'Billie Eilish',
      tour: 'HIT ME HARD AND SOFT',
      img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
      reviewText: 'Lorem ipsum odor amet, consectetuer adipiscing elit. Potenti mus fermentum sed dapibus egestas; aptent faucibus quisque? Fames fringilla consectetur tortor leo potenti at porttitor aenean.',
      attachedImg: null,
      setlist: null,
      rating: 5,
      likes: 10,
      comments: 3,
    },
    {
      type: 'wishlist',
      username: 'Jane Smith',
      avatar: 'imgs/user-profile.jpeg',
      postDate: 'Mar 25, 2025',
      date: null,
      venue: null,
      artist: 'Bad Bunny',
      tour: 'No me quiero ir de aqui',
      img: 'https://i.scdn.co/image/ab6761610000e5eb81f47f44084e0a09b5f0fa13',
      reviewText: null,
      attachedImg: null,
      setlist: null,
      rating: null,
      likes: 2,
      comments: 0,
    }
  ];

  followers: Followers[] = [
    {
      userName: "pinky234"
    },
    {
      userName: "reddy123"
    }
  ];

  following: Following[] = [
    {
      userName: "mrbrown234"
    },
    {
      userName: "heart123"
    }
  ]

  // Convert concert cards to post format
  convertToPost(concertCard: ConcertCard, type: string = 'review'): Post {
    return {
      type: type,
      username: this.userProfile.name,
      avatar: this.userProfile.profileImage,
      postDate: 'Mar 30, 2025',
      date: concertCard.date,
      venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
      artist: concertCard.artist,
      tour: concertCard.title,
      img: concertCard.image,
      reviewText: null,
      attachedImg: null,
      setlist: null,
      rating: 5,
      likes: 10,
      comments: 3,
    };
  }

  constructor() {}

  getUserProfile(): Observable<UserProfile> {
    return of(this.userProfile);
  }

  getFavoriteConcerts(): Observable<ConcertCard[]> {
    return of(this.favoriteConcerts);
  }

  getRecentAttendance(): Observable<ConcertCard[]> {
    return of(this.recentAttendance);
  }

  getBucketList(): Observable<ConcertCard[]> {
    return of(this.bucketList);
  }

  getRecentActivity(): Observable<Activity[]> {
    return of(this.recentActivity);
  }

  getRecentLists(): Observable<List[]> {
    return of(this.recentLists);
  }

  // Get user posts 
  getUserPosts(): Observable<Post[]> {
    return of(this.userPosts);
  }

  // Get filtered posts by type
  getPostsByType(type: string): Observable<Post[]> {
    return of(this.userPosts.filter(post => post.type === type));
  }
}