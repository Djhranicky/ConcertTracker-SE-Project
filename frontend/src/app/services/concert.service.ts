import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { of } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Concert, Artist } from '../models/artist.model';
@Injectable({
  providedIn: 'root',
})
export class ConcertService {
  concert1: Concert = {
    date: 'Feb 19, 2025',
    venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
    setlist: `[\n    {\n        \"name\": \"\",\n        \"tape\": true,\n        \"info\": \"contains elements of “THE GREATEST\\\"\"\n    },\n    {\n        \"name\": \"CHIHIRO\"\n    },\n    {\n        \"name\": \"LUNCH\"\n    },\n    {\n        \"name\": \"NDA\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"Therefore I Am\"\n    },\n    {\n        \"name\": \"WILDFLOWER\"\n    },\n    {\n        \"name\": \"when the party's over\",\n        \"info\": \"live looped vocal layers intro\"\n    },\n    {\n        \"name\": \"THE DINER\"\n    },\n    {\n        \"name\": \"ilomilo\"\n    },\n    {\n        \"name\": \"bad guy\"\n    },\n    {\n        \"name\": \"THE GREATEST\"\n    },\n    {\n        \"name\": \"Your Power\"\n    },\n    {\n        \"name\": \"SKINNY\"\n    },\n    {\n        \"name\": \"TV\"\n    },\n    {\n        \"name\": \"BITTERSUITE\",\n        \"tape\": true,\n        \"info\": \"Transition\"\n    },\n    {\n        \"name\": \"bury a friend\"\n    },\n    {\n        \"name\": \"Oxytocin\"\n    },\n    {\n        \"name\": \"Guess\",\n        \"cover\": {\n            \"mbid\": \"260b6184-8828-48eb-945c-bc4cb6fc34ca\",\n            \"name\": \"Charli xcx\",\n            \"sortName\": \"Charli xcx\",\n            \"disambiguation\": \"\",\n            \"url\": \"https:\/\/www.setlist.fm\/setlists\/charli-xcx-33d5dcdd.html\"\n        },\n        \"info\": \"featuring Billie Eilish\"\n    },\n    {\n        \"name\": \"everything i wanted\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"TRUE BLUE\\\" version\"\n    },\n    {\n        \"name\": \"lovely \/ BLUE \/ ocean eyes\",\n        \"info\": \"\\\"BORN BLUE\\\" version performed\"\n    },\n    {\n        \"name\": \"L'AMOUR DE MA VIE\",\n        \"info\": \"with \\\"OVER NOW EXTENDED EDIT\\\"\"\n    },\n    {\n        \"name\": \"What Was I Made For?\"\n    },\n    {\n        \"name\": \"Happier Than Ever\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"BIRDS OF A FEATHER\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"BORN BLUE\\\" version\"\n    }\n]`,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  concert2: Concert = {
    date: 'Feb 21, 2025',
    venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
    setlist: `[\n    {\n        \"name\": \"\",\n        \"tape\": true,\n        \"info\": \"contains elements of “THE GREATEST\\\"\"\n    },\n    {\n        \"name\": \"CHIHIRO\"\n    },\n    {\n        \"name\": \"LUNCH\"\n    },\n    {\n        \"name\": \"NDA\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"Therefore I Am\"\n    },\n    {\n        \"name\": \"WILDFLOWER\"\n    },\n    {\n        \"name\": \"when the party's over\",\n        \"info\": \"live looped vocal layers intro\"\n    },\n    {\n        \"name\": \"THE DINER\"\n    },\n    {\n        \"name\": \"ilomilo\"\n    },\n    {\n        \"name\": \"bad guy\"\n    },\n    {\n        \"name\": \"THE GREATEST\"\n    },\n    {\n        \"name\": \"Your Power\"\n    },\n    {\n        \"name\": \"SKINNY\"\n    },\n    {\n        \"name\": \"TV\"\n    },\n    {\n        \"name\": \"BITTERSUITE\",\n        \"tape\": true,\n        \"info\": \"Transition\"\n    },\n    {\n        \"name\": \"bury a friend\"\n    },\n    {\n        \"name\": \"Oxytocin\"\n    },\n    {\n        \"name\": \"Guess\",\n        \"cover\": {\n            \"mbid\": \"260b6184-8828-48eb-945c-bc4cb6fc34ca\",\n            \"name\": \"Charli xcx\",\n            \"sortName\": \"Charli xcx\",\n            \"disambiguation\": \"\",\n            \"url\": \"https:\/\/www.setlist.fm\/setlists\/charli-xcx-33d5dcdd.html\"\n        },\n        \"info\": \"featuring Billie Eilish\"\n    },\n    {\n        \"name\": \"everything i wanted\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"TRUE BLUE\\\" version\"\n    },\n    {\n        \"name\": \"lovely \/ BLUE \/ ocean eyes\",\n        \"info\": \"\\\"BORN BLUE\\\" version performed\"\n    },\n    {\n        \"name\": \"L'AMOUR DE MA VIE\",\n        \"info\": \"with \\\"OVER NOW EXTENDED EDIT\\\"\"\n    },\n    {\n        \"name\": \"What Was I Made For?\"\n    },\n    {\n        \"name\": \"Happier Than Ever\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"BIRDS OF A FEATHER\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"BORN BLUE\\\" version\"\n    }\n]`,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  concert3: Concert = {
    date: 'Feb 22, 2025',
    venue: 'Brisbane Entertainment Centre, Brisbane, Australia',
    setlist: `[\n    {\n        \"name\": \"\",\n        \"tape\": true,\n        \"info\": \"contains elements of “THE GREATEST\\\"\"\n    },\n    {\n        \"name\": \"CHIHIRO\"\n    },\n    {\n        \"name\": \"LUNCH\"\n    },\n    {\n        \"name\": \"NDA\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"Therefore I Am\"\n    },\n    {\n        \"name\": \"WILDFLOWER\"\n    },\n    {\n        \"name\": \"when the party's over\",\n        \"info\": \"live looped vocal layers intro\"\n    },\n    {\n        \"name\": \"THE DINER\"\n    },\n    {\n        \"name\": \"ilomilo\"\n    },\n    {\n        \"name\": \"bad guy\"\n    },\n    {\n        \"name\": \"THE GREATEST\"\n    },\n    {\n        \"name\": \"Your Power\"\n    },\n    {\n        \"name\": \"SKINNY\"\n    },\n    {\n        \"name\": \"TV\"\n    },\n    {\n        \"name\": \"BITTERSUITE\",\n        \"tape\": true,\n        \"info\": \"Transition\"\n    },\n    {\n        \"name\": \"bury a friend\"\n    },\n    {\n        \"name\": \"Oxytocin\"\n    },\n    {\n        \"name\": \"Guess\",\n        \"cover\": {\n            \"mbid\": \"260b6184-8828-48eb-945c-bc4cb6fc34ca\",\n            \"name\": \"Charli xcx\",\n            \"sortName\": \"Charli xcx\",\n            \"disambiguation\": \"\",\n            \"url\": \"https:\/\/www.setlist.fm\/setlists\/charli-xcx-33d5dcdd.html\"\n        },\n        \"info\": \"featuring Billie Eilish\"\n    },\n    {\n        \"name\": \"everything i wanted\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"TRUE BLUE\\\" version\"\n    },\n    {\n        \"name\": \"lovely \/ BLUE \/ ocean eyes\",\n        \"info\": \"\\\"BORN BLUE\\\" version performed\"\n    },\n    {\n        \"name\": \"L'AMOUR DE MA VIE\",\n        \"info\": \"with \\\"OVER NOW EXTENDED EDIT\\\"\"\n    },\n    {\n        \"name\": \"What Was I Made For?\"\n    },\n    {\n        \"name\": \"Happier Than Ever\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"BIRDS OF A FEATHER\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"BORN BLUE\\\" version\"\n    }\n]`,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  concert4: Concert = {
    date: 'Feb 24, 2025',
    venue: 'Qudos Bank Arena, Sydney, Australia',
    setlist: `[\n    {\n        \"name\": \"\",\n        \"tape\": true,\n        \"info\": \"contains elements of “THE GREATEST\\\"\"\n    },\n    {\n        \"name\": \"CHIHIRO\"\n    },\n    {\n        \"name\": \"LUNCH\"\n    },\n    {\n        \"name\": \"NDA\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"Therefore I Am\"\n    },\n    {\n        \"name\": \"WILDFLOWER\"\n    },\n    {\n        \"name\": \"when the party's over\",\n        \"info\": \"live looped vocal layers intro\"\n    },\n    {\n        \"name\": \"THE DINER\"\n    },\n    {\n        \"name\": \"ilomilo\"\n    },\n    {\n        \"name\": \"bad guy\"\n    },\n    {\n        \"name\": \"THE GREATEST\"\n    },\n    {\n        \"name\": \"Your Power\"\n    },\n    {\n        \"name\": \"SKINNY\"\n    },\n    {\n        \"name\": \"TV\"\n    },\n    {\n        \"name\": \"BITTERSUITE\",\n        \"tape\": true,\n        \"info\": \"Transition\"\n    },\n    {\n        \"name\": \"bury a friend\"\n    },\n    {\n        \"name\": \"Oxytocin\"\n    },\n    {\n        \"name\": \"Guess\",\n        \"cover\": {\n            \"mbid\": \"260b6184-8828-48eb-945c-bc4cb6fc34ca\",\n            \"name\": \"Charli xcx\",\n            \"sortName\": \"Charli xcx\",\n            \"disambiguation\": \"\",\n            \"url\": \"https:\/\/www.setlist.fm\/setlists\/charli-xcx-33d5dcdd.html\"\n        },\n        \"info\": \"featuring Billie Eilish\"\n    },\n    {\n        \"name\": \"everything i wanted\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"TRUE BLUE\\\" version\"\n    },\n    {\n        \"name\": \"lovely \/ BLUE \/ ocean eyes\",\n        \"info\": \"\\\"BORN BLUE\\\" version performed\"\n    },\n    {\n        \"name\": \"L'AMOUR DE MA VIE\",\n        \"info\": \"with \\\"OVER NOW EXTENDED EDIT\\\"\"\n    },\n    {\n        \"name\": \"What Was I Made For?\"\n    },\n    {\n        \"name\": \"Happier Than Ever\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"BIRDS OF A FEATHER\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"BORN BLUE\\\" version\"\n    }\n]`,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  concert5: Concert = {
    date: 'Feb 25, 2025',
    venue: 'Qudos Bank Arena, Sydney, Australia',
    setlist: `[\n    {\n        \"name\": \"\",\n        \"tape\": true,\n        \"info\": \"contains elements of “THE GREATEST\\\"\"\n    },\n    {\n        \"name\": \"CHIHIRO\"\n    },\n    {\n        \"name\": \"LUNCH\"\n    },\n    {\n        \"name\": \"NDA\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"Therefore I Am\"\n    },\n    {\n        \"name\": \"WILDFLOWER\"\n    },\n    {\n        \"name\": \"when the party's over\",\n        \"info\": \"live looped vocal layers intro\"\n    },\n    {\n        \"name\": \"THE DINER\"\n    },\n    {\n        \"name\": \"ilomilo\"\n    },\n    {\n        \"name\": \"bad guy\"\n    },\n    {\n        \"name\": \"THE GREATEST\"\n    },\n    {\n        \"name\": \"Your Power\"\n    },\n    {\n        \"name\": \"SKINNY\"\n    },\n    {\n        \"name\": \"TV\"\n    },\n    {\n        \"name\": \"BITTERSUITE\",\n        \"tape\": true,\n        \"info\": \"Transition\"\n    },\n    {\n        \"name\": \"bury a friend\"\n    },\n    {\n        \"name\": \"Oxytocin\"\n    },\n    {\n        \"name\": \"Guess\",\n        \"cover\": {\n            \"mbid\": \"260b6184-8828-48eb-945c-bc4cb6fc34ca\",\n            \"name\": \"Charli xcx\",\n            \"sortName\": \"Charli xcx\",\n            \"disambiguation\": \"\",\n            \"url\": \"https:\/\/www.setlist.fm\/setlists\/charli-xcx-33d5dcdd.html\"\n        },\n        \"info\": \"featuring Billie Eilish\"\n    },\n    {\n        \"name\": \"everything i wanted\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"TRUE BLUE\\\" version\"\n    },\n    {\n        \"name\": \"lovely \/ BLUE \/ ocean eyes\",\n        \"info\": \"\\\"BORN BLUE\\\" version performed\"\n    },\n    {\n        \"name\": \"L'AMOUR DE MA VIE\",\n        \"info\": \"with \\\"OVER NOW EXTENDED EDIT\\\"\"\n    },\n    {\n        \"name\": \"What Was I Made For?\"\n    },\n    {\n        \"name\": \"Happier Than Ever\",\n        \"info\": \"Shortened\"\n    },\n    {\n        \"name\": \"BIRDS OF A FEATHER\"\n    },\n    {\n        \"name\": \"BLUE\",\n        \"tape\": true,\n        \"info\": \"\\\"BORN BLUE\\\" version\"\n    }\n]`,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  artist1: Artist = {
    name: 'Billie Eilish',
    imageUrl:
      'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',

    ID: 0,
    MBID: 'mbid',
    lastUpdate: new Date('2025-03-31T21:30:48.0945199-04:00'),
    recentSetlists: null,
    toursCount: 20,
    showsCount: 150,
    upcomingShows: null,
  };

  upcoming1: Concert = {
    date: 'Apr 23, 2025',
    venue: 'Avicii Arena, Stockholm, Sweden',
    setlist: null,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  upcoming2: Concert = {
    date: 'Apr 24, 2025',
    venue: 'Avicii Arena, Stockholm, Sweden',
    setlist: null,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  upcoming3: Concert = {
    date: 'Apr 26, 2025',
    venue: 'Unity Arena, Fornebu, Norway',
    setlist: null,
    tour: 'HIT ME HARD AND SOFT',
    artist: 'Billie Eilish',
    img: 'https://res.cloudinary.com/hits-photos-archive/image/upload/v1736890770/legacy-migration/legacy-hitsdd_photo_gal__photo_1891402125.png',
  };

  data = {
    concerts: [
      this.concert5,
      this.concert4,
      this.concert3,
      this.concert2,
      this.concert1,
    ],
    upcoming: [this.upcoming1, this.upcoming2, this.upcoming3],
  };

  private url = 'http://localhost:8080/api';
  constructor(private http: HttpClient) {}

  getConcert(): Concert {
    return this.concert1;
  }

  getArtist(): Artist {
    return this.artist1;
  }

  getArtistByName(name: string): Observable<Artist> {
    const encoded = encodeURIComponent(name);
    return this.http.get<any>(`${this.url}/artist?name=${encoded}`).pipe(
      map((response: any) => ({
        ID: response.artist.ID,
        MBID: response.artist.MBID,
        name: response.artist.name,
        lastUpdate: response.artist.UpdatedAt,
        imageUrl: null,
        recentSetlists: null,
        toursCount: response.number_of_tours,
        showsCount: response.total_setlists,
        upcomingShows: null,
      }))
    );
  }

  getRecentConcerts(): Observable<Concert[]> {
    return of(this.data.concerts);
  }

  getUpcomingConcerts(): Observable<Concert[]> {
    return of(this.data.upcoming);
  }
}
