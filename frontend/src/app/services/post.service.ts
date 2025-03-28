import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { of } from 'rxjs';
import { Tour } from './pop-tours.service';

export interface User {
  username: string;
  avatar: string;
}

export interface Concert extends Tour {
  date: string | null;
  venue: string | null;
}

export interface Post extends User, Concert, Tour {
  type: string;
  postDate: string;
  reviewText: string | null;
  attachedImg: string | null;
  rating: number | null;
  likes: number;
  comments: number;
}

@Injectable({
  providedIn: 'root',
})
export class PostService {
  post1: Post = {
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
    rating: 4,
    likes: 10,
    comments: 3,
  };

  post2: Post = {
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
    likes: 2,
    comments: 0,
  };

  post3: Post = {
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
    rating: 5,
    likes: 2,
    comments: 0,
  };

  data = {
    posts: [this.post1, this.post2, this.post3],
  };
  constructor(private http: HttpClient) {}

  getPosts(): Observable<Post[]> {
    return of(this.data.posts);
  }
}
