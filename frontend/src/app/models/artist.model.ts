export interface Artist {
  ID: number;
  MBID: string;
  name: string;
  lastUpdate: Date;
  imageUrl: string | null;
  recentSetlists: Concert[] | null;
  toursCount: number;
  showsCount: number;
  upcomingShows: Concert[] | null;
}

export interface Tour {
  artist: string;
  img: string | null;
  tour: string;
}

export interface Concert extends Tour {
  city: string | null;
  date: string | null;
  id: string | null;
  venue: string | null;
  setlist: Song[] | null;
}
export interface Song {
  name: string | null;
  with: string | null;
  order: number;
  info: string | null;
  tape: boolean;
  cover: Cover | null;
}
export interface Cover {
  mbid: string | null;
  name: string | null;
}
