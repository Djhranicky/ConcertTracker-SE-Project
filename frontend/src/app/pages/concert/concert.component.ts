import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { Card } from 'primeng/card';
import { AvatarModule } from 'primeng/avatar';
import { AvatarGroup } from 'primeng/avatargroup';
import { ImageModule } from 'primeng/image';
import { Button } from 'primeng/button';
import { Timeline } from 'primeng/timeline';
import { Concert, Song } from '../../models/artist.model';
import { ConcertService } from '../../services/concert.service';
import { Post, PostService } from '../../services/post.service';
import { ActivatedRoute } from '@angular/router';
import { FriendlyDatePipe } from '../../utils/friendlyDate.pipe';
import { ProgressSpinner } from 'primeng/progressspinner';
import { RouterModule } from '@angular/router';
@Component({
  selector: 'app-concert',
  imports: [
    NavbarComponent,
    Card,
    ImageModule,
    Button,
    CommonModule,
    AvatarModule,
    AvatarGroup,
    Timeline,
    ProgressSpinner,
    FriendlyDatePipe,
    RouterModule,
  ],
  templateUrl: './concert.component.html',
  styleUrl: './concert.component.css',
  providers: [ConcertService, PostService],
})
export class ConcertComponent {
  @Input() concert: Concert;
  posts: Post[] = [];
  day: string;
  month: string;
  year: string;
  loading: boolean = true;
  id: string;

  constructor(
    private concertService: ConcertService,
    private postService: PostService,
    private route: ActivatedRoute
  ) {}

  // parseSetlist() {
  //   if (this.concert.setlist) {
  //     // this.setlist = JSON.parse(this.concert.setlist);
  //     // console.log(this.setlist);
  //   }
  // }

  objectEntries(obj: any): [string, any][] {
    return Object.entries(obj);
  }

  ngOnInit() {
    this.route.paramMap.subscribe((params) => {
      this.id = params.get('id') as string;
      this.loading = true;
    });

    this.concertService.getConcert(this.id).subscribe((concert) => {
      this.concert = concert;
      console.log(this.concert);
      this.toggleLoading();
    });

    // this.parseSetlist();

    this.postService.getPosts().subscribe((data) => {
      this.posts = data;
      // console.log(this.posts);
    });

    // console.log(this.concert);
  }

  private toggleLoading() {
    if (this.loading) this.loading = false;
  }
}
