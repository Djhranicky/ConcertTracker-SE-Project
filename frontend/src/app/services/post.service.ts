import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { of } from 'rxjs';
import { Post } from '../models/post.model';

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
    venue: 'Brisbane Entertainment Centre',
    artist: 'Billie Eilish',
    setlist: null,
    tour: 'HIT ME HARD AND SOFT',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
    reviewText: null,
    attachedImg: null,
    rating: 4,
    likes: 10,
    comments: 3,
    city: 'Brisbane, Australia',
    id: '',
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
    setlist: null,
    reviewText: null,
    attachedImg: null,
    rating: 4,
    likes: 2,
    comments: 0,
    city: '',
    id: '',
  };

  post3: Post = {
    type: 'review',
    username: 'Jane Doe',
    avatar: 'imgurl',
    postDate: 'Feb 20, 2025',
    date: 'Feb 19, 2025',
    venue: 'Brisbane Entertainment Centre',
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
    city: 'Brisbane, Australia',
    id: '',
  };

  data = {
    posts: [this.post1, this.post2, this.post3],
  };

  private url = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  getPosts(): Observable<Post[]> {
    return of(this.data.posts);
  }

  postPost(payload: {
    authorUsername: string;
    externalConcertID: string;
    isPublic: boolean;
    rating: number | null;
    text: string | null;
    type: string; //'ATTENDED' | 'WISHLIST' | 'REVIEW' | 'LISTCREATED';
    // userPostID: number | null;
  }) {
    return this.http.post(`${this.url}/userpost`, payload);
  }

  getDashboardPosts(username: string, p: number = 1): Observable<Post[]> {
    const params = new HttpParams().set('username', username);

    return this.http.get<any[]>(`${this.url}/userpost`, { params }).pipe(
      map((response) =>
        response.map(
          (item): Post => ({
            // Fields from Post interface
            type: item.type,
            postDate: item.createdAt,
            reviewText: item.text ?? null,
            attachedImg: null,
            rating: item.rating ?? null,
            likes: 0,
            comments: 0,
            avatar: '',
            id: item.postID,

            username: item.authorUsername,

            date: item.concertDate,
            venue: item.venueName,
            city: item.venueCity,

            tour: item.tourName,
            artist: item.artistName,
            img: null,
            setlist: null,
          })
        )
      )
    );
  }
}
