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
import { Post } from '../../models/post.model';
import { PostService } from '../../services/post.service';
import { ActivatedRoute } from '@angular/router';
import { FriendlyDatePipe } from '../../utils/friendlyDate.pipe';
import { ProgressSpinner } from 'primeng/progressspinner';
import { RouterModule } from '@angular/router';

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
    FriendlyDatePipe,
    ProgressSpinner,
    RouterModule,
  ],
  templateUrl: './artist.component.html',
  styleUrl: './artist.component.css',
  providers: [ConcertService, PostService],
})
export class ArtistComponent implements OnInit {
  @Input() artist: Artist;
  name: string;
  upcoming: Concert[] = [];
  responsiveOptions: any[] | undefined;
  upcomingTuples: [Concert, string[]][] = [];
  mostPlayed: [string, string][] = [];
  loading: boolean = true;

  constructor(
    private concertService: ConcertService,
    private postService: PostService,
    private route: ActivatedRoute
  ) {}

  ngOnInit() {
    this.route.paramMap.subscribe((params) => {
      this.name = params.get('name') as string;

      this.loading = true;
    });

    this.concertService.getArtist(this.name).subscribe((artist) => {
      this.artist = artist;
      //console.log(this.artist);
      this.toggleLoading();
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

  private toggleLoading() {
    if (this.loading) this.loading = false;
  }
}
