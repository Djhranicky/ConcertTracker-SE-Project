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

export interface Following {
  username: string;
}

export interface Followers {
  username: string;
}