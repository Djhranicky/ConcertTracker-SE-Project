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
  loadingCount = 0;
  stats: [string, unknown][];

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
      console.log(this.artist);
      this.toggleLoading();
    });

    this.concertService.getUpcomingConcerts(this.name).subscribe((upcoming) => {
      this.upcoming = upcoming;
      this.toggleLoading();
    });

    this.concertService.getStats(this.name).subscribe((stats) => {
      this.stats = Object.entries(stats);
      console.log('stats', this.stats);
      this.toggleLoading();
    });

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
    this.loadingCount++;
    if (this.loadingCount === 3) this.loading = false;
  }
}
