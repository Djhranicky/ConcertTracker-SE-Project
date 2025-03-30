import { Component, Input, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PostService, Post } from '../../services/post.service';
import { UserService } from '../../services/user.service';
import { Button } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { Avatar } from 'primeng/avatar';
import { TimeAgoPipe } from '../../utils/time-ago.pipe';

@Component({
  selector: 'app-post',
  imports: [CardModule, Button, Avatar, CommonModule, TimeAgoPipe],
  templateUrl: './post.component.html',
  styleUrl: './post.component.css',
  providers: [PostService, UserService],
})
export class PostComponent implements OnInit {
  isLiked: boolean = false;
  isCurrentUser: boolean = false;
  currentUsername: string = '';

  @Input() post: Post;
  
  constructor(private userService: UserService) {}
  
  ngOnInit(): void {
    // Check if the post is from the current user
    this.userService.getUserProfile().subscribe(profile => {
      this.currentUsername = profile.name;
      this.isCurrentUser = this.post.username === this.currentUsername;
    });
  }

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
  
  getDisplayUsername(): string {
    return this.isCurrentUser ? 'You' : this.post.username;
  }
}