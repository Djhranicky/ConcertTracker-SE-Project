import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PostComponent } from '../../components/post/post.component';
import { PostService } from '../../services/post.service';
import { Post } from '../../models/post.model';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { Button } from 'primeng/button';
import { Card } from 'primeng/card';
import { AvatarModule } from 'primeng/avatar';

@Component({
  selector: 'app-dashboard',
  imports: [
    PostComponent,
    CommonModule,
    NavbarComponent,
    Button,
    Card,
    AvatarModule,
  ],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
  providers: [PostService],
})
export class DashboardComponent implements OnInit {
  posts: Post[] = [];

  constructor(private postsService: PostService) {}

  ngOnInit(): void {
    this.postsService.getPosts().subscribe((data) => {
      this.posts = data;
      // console.log(this.posts);
    });
  }
}
