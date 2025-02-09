import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { of } from 'rxjs';

export interface Tour {
  artist: string;
  tour: string;
  img: string;
}

@Injectable({
  providedIn: 'root',
})
export class PopToursService {
  data = {
    tours: [
      {
        artist: 'Coldplay',
        tour: 'Music of the Spheres',
        img: 'https://seatgeekimages.com/performers-landscape/coldplay-827fc3/32/1100x1900.jpg?auto=webp&width=3840&quality=75',
      },
      {
        artist: 'Tyler, the Creator',
        tour: '30 Minutes of Chromakopia',
        img: 'https://static01.nyt.com/images/2021/07/08/arts/08tyler-review2/merlin_190548804_d5cb859f-7f6b-4de0-a078-afd60438d478-articleLarge.jpg?quality=75&auto=webp&disable=upscale',
      },
      {
        artist: 'Oasis',
        tour: "Live '25",
        img: 'https://relix.com/wp-content/uploads/2024/08/unnamed-25-1.jpg',
      },
      {
        artist: 'Kendrick Lamar',
        tour: 'Grand National Tour',
        img: 'https://mediaproxy.tvtropes.org/width/1200/https://static.tvtropes.org/pmwiki/pub/images/kendricklamar.png',
      },
      {
        artist: 'Shakira',
        tour: 'Las Mujeres Ya No Lloran',
        img: 'https://upload.wikimedia.org/wikipedia/commons/b/b8/2023-11-16_Gala_de_los_Latin_Grammy%2C_03_%28cropped%2902.jpg',
      },
    ],
  };
  constructor(private http: HttpClient) {}

  getPopTours(): Observable<Tour[]> {
    return of(this.data.tours);
  }
}
