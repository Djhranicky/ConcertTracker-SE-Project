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
  img: string;
  tour: string;
}

export interface Concert extends Tour {
  date: string | null;
  venue: string | null;
  setlist: string | null;
}
export interface Song {
  name: string | null;
  with: string | null;
  cover: Cover | null;
  info: string | null;
  tape: boolean;
}
export interface Cover {
  mbid: string | null;
  name: string | null;
  sortName: string | null;
  disambiguation: string | null;
  url: string | null;
}
