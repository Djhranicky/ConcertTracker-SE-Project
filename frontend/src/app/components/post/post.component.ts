import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PostService, Post } from '../../services/post.service';
import { Button } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { Avatar } from 'primeng/avatar';
import { TimeAgoPipe } from '../../utils/time-ago.pipe';
@Component({
  selector: 'app-post',
  imports: [CardModule, Button, Avatar, CommonModule, TimeAgoPipe],
  templateUrl: './post.component.html',
  styleUrl: './post.component.css',
  providers: [PostService],
})
export class PostComponent {
  isLiked: boolean = false;

  @Input() post: Post;
  // responsiveOptions: any[] | undefined;

  getStars(rating: number): number[] {
    return Array(rating).fill(0);
  }

  toggleLike() {
    this.isLiked = !this.isLiked;
    if (this.isLiked) {
      this.post.likes = this.post.likes + 1;
    } else {
      this.post.likes = this.post.likes - 1;
    }
  }
}
