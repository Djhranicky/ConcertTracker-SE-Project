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
import { PostService } from '../../services/post.service';
import { ActivatedRoute } from '@angular/router';
import { FriendlyDatePipe } from '../../utils/friendlyDate.pipe';
import { ProgressSpinner } from 'primeng/progressspinner';
import { RouterModule } from '@angular/router';
import { Post } from '../../models/post.model';
import { DialogModule } from 'primeng/dialog';
import { TextareaModule } from 'primeng/textarea';
import { Checkbox } from 'primeng/checkbox';
import { RatingModule } from 'primeng/rating';
import { ToggleButtonModule } from 'primeng/togglebutton';
import { FormsModule } from '@angular/forms';
import { FloatLabel } from 'primeng/floatlabel';
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
    DialogModule,
    TextareaModule,
    Checkbox,
    RatingModule,
    ToggleButtonModule,
    FormsModule,
    FloatLabel,
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
  showModal: boolean = false;
  reviewText: string = '';
  rating: number = 0;
  isPublic: boolean = true;
  type: boolean;
  user: string;
  userID: number;

  constructor(
    private concertService: ConcertService,
    private postService: PostService,
    private route: ActivatedRoute
  ) {}

  objectEntries(obj: any): [string, any][] {
    return Object.entries(obj);
  }

  ngOnInit() {
    this.user = localStorage.getItem('user') as string;
    this.userID = Number(localStorage.getItem('id') as string);

    this.route.paramMap.subscribe((params) => {
      this.id = params.get('id') as string;
      this.loading = true;
    });

    this.concertService.getConcert(this.id).subscribe((concert) => {
      this.concert = concert;
      console.log(this.concert);
      this.toggleLoading();
    });

    this.postService.getPosts().subscribe((data) => {
      this.posts = data;
    });
  }

  private toggleLoading() {
    if (this.loading) this.loading = false;
  }

  openModal() {
    this.showModal = true;
  }

  submitPost() {
    let postType = '';
    if (this.type == true) {
      postType = 'ATTENDED';
    } else {
      postType = 'WISHLIST';
    }

    const payload = {
      authorID: this.userID,
      concertID: this.id, // Also assign dynamically based on the concert context
      isPublic: this.isPublic,
      rating: this.rating,
      text: this.reviewText,
      type: postType,
    };
    this.postService.postPost(payload).subscribe({
      next: () => {
        console.log('Post submitted');
        this.showModal = false;
      },
      error: (err) => {
        console.error('Failed to submit post', err);
      },
    });
  }
}
