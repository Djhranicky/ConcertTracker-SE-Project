import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PostComponent } from '../../components/post/post.component';
import { PostService, Post } from '../../services/post.service';
import { NavbarComponent } from '../../components/navbar/navbar.component';
@Component({
  selector: 'app-dashboard',
  imports: [PostComponent, CommonModule, NavbarComponent],
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
      console.log(this.posts);
    });
  }
}
