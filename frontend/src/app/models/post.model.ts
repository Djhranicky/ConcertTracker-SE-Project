import { Concert, Tour } from './artist.model';
export interface User {
  username: string;
  avatar: string;
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
