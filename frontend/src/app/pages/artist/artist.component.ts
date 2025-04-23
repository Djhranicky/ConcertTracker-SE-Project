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
import { Button } from 'primeng/button';
import { Location } from '@angular/common';
import { forkJoin, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

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
    Button,
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
  loadingTarget = 1;
  stats: [string, unknown][];
  mbid: string;
  recent: (string | null)[] | undefined;
  tours: string[] = [];

  constructor(
    private concertService: ConcertService,
    private postService: PostService,
    private location: Location,
    private route: ActivatedRoute
  ) {}

  ngOnInit() {
    this.route.paramMap.subscribe((params) => {
      this.name = params.get('name') as string;
      this.loading = true;
    });

    this.concertService.getArtist(this.name).subscribe((artist) => {
      this.artist = artist;
      this.mbid = artist.MBID;
      const recentIds = artist.recentSetlists?.map((item) => item.id);
      console.log('recent', recentIds);

      if (recentIds && recentIds.length > 0) {
        const concertObservables = recentIds.map((id) => {
          if (!id) return of('');
          return this.concertService.getConcert(id).pipe(
            map((concert) => concert?.tour || ''),
            catchError(() => of(''))
          );
        });

        forkJoin(concertObservables).subscribe((tourList) => {
          this.tours = tourList;

          // Now update each recentSetlist with its corresponding tour
          this.artist.recentSetlists?.forEach((setlist, index) => {
            setlist.tour = this.tours[index];
          });

          console.log(this.artist, this.tours);
          this.toggleLoading();
        });
      } else {
        this.toggleLoading();
      }
    });

    if (
      !(this.artist?.showsCount === 20 && this.artist?.upcomingShows === null)
    ) {
      this.loadingTarget = 3;
      this.concertService
        .getUpcomingConcerts(this.name)
        .subscribe((upcoming) => {
          this.upcoming = upcoming;
          this.toggleLoading();
        });

      this.concertService.getStats(this.name).subscribe((stats) => {
        if (stats) {
          this.stats = Object.entries(stats);
          console.log('stats', this.stats);
        }
        this.toggleLoading();
      });
    }

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

  runImport() {
    if (!this.mbid) return;

    this.concertService.runFullImport(this.mbid).subscribe({
      next: () => window.location.reload(),
      error: (err) => console.error('Import failed:', err),
    });
  }

  private toggleLoading() {
    this.loadingCount++;
    console.log(this.loadingCount, this.loadingTarget);
    if (this.loadingCount === this.loadingTarget) this.loading = false;
  }
}
