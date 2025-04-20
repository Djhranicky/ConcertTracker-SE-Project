import { Component, Input, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { Card } from 'primeng/card';
import { AvatarModule } from 'primeng/avatar';
import { ImageModule } from 'primeng/image';
import { Carousel, CarouselModule } from 'primeng/carousel';
import { Tag } from 'primeng/tag';
import { TableModule } from 'primeng/table';
import { ConcertService } from '../../services/concert.service';
import { Artist, Concert, Song } from '../../models/artist.model';
import { Post, PostService } from '../../services/post.service';

@Component({
  selector: 'app-artist',
  imports: [
    NavbarComponent,
    Card,
    ImageModule,
    CommonModule,
    AvatarModule,
    CarouselModule,
    Tag,
    TableModule,
  ],
  templateUrl: './artist.component.html',
  styleUrl: './artist.component.css',
  providers: [ConcertService, PostService],
})
export class ArtistComponent implements OnInit {
  @Input() artist: Artist;

  concerts: Concert[] = [];
  upcoming: Concert[] = [];
  responsiveOptions: any[] | undefined;
  upcomingTuples: [Concert, string[]][] = [];
  mostPlayed: [string, string][] = [];

  constructor(
    private concertService: ConcertService,
    private postService: PostService
  ) {}

  ngOnInit() {
    this.artist = this.concertService.getArtist();
    this.concertService.getRecentConcerts().subscribe((concerts) => {
      this.concerts = concerts;
    });
    this.concertService.getUpcomingConcerts().subscribe((upcoming) => {
      this.upcoming = upcoming;
    });
    for (const item of this.upcoming) {
      let date = item.date!.split(' ');
      let month = date[0];
      let day = date[1].slice(0, -1);
      let year = date[2];
      this.upcomingTuples.push([item, [month, day, year]]);
    }
    this.mostPlayed = [
      ["when the party's over", '(315)'],
      ['ocean eyes', '(274)'],
      ['bellyache', '(268)'],
    ];

    // console.log(this.upcomingTuples);

    this.responsiveOptions = [
      {
        breakpoint: '1400px',
        numVisible: 5,
        numScroll: 1,
      },
      {
        breakpoint: '1199px',
        numVisible: 3,
        numScroll: 1,
      },
      {
        breakpoint: '850px',
        numVisible: 2,
        numScroll: 1,
      },
      {
        breakpoint: '575px',
        numVisible: 1,
        numScroll: 1,
      },
    ];
  }
}
