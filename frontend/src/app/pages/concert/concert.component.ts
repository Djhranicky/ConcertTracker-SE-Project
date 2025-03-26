import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { Card } from 'primeng/card';
import { AvatarModule } from 'primeng/avatar';
import { ImageModule } from 'primeng/image';
import { Button } from 'primeng/button';
import { Timeline } from 'primeng/timeline';
import { Concert, Song, ConcertService } from '../../services/concert.service';
import { Post, PostService } from '../../services/post.service';

@Component({
  selector: 'app-concert',
  imports: [
    NavbarComponent,
    Card,
    ImageModule,
    Button,
    CommonModule,
    AvatarModule,
    Timeline,
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
  setlist: Song[];
  constructor(
    private concertService: ConcertService,
    private postService: PostService
  ) {}

  parseSetlist() {
    if (this.concert.setlist) {
      this.setlist = JSON.parse(this.concert.setlist);
      console.log(this.setlist);
    }
  }

  objectEntries(obj: any): [string, any][] {
    return Object.entries(obj);
  }

  ngOnInit() {
    this.concert = this.concertService.getConcert();
    let date = this.concert.date!.split(' ');
    this.month = date[0];
    this.day = date[1].slice(0, -1);
    this.year = date[2];
    this.parseSetlist();
    this.postService.getPosts().subscribe((data) => {
      this.posts = data;
      console.log(this.posts);
    });

    // console.log(this.concert);
  }
}
