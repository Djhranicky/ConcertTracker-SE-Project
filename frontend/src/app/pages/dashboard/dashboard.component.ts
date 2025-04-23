import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PostComponent } from '../../components/post/post.component';
import { PostService } from '../../services/post.service';
import { Post } from '../../models/post.model';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { Button } from 'primeng/button';
import { Card } from 'primeng/card';
import { AvatarModule } from 'primeng/avatar';
import { UserService } from '../../services/user.service';
import { RouterModule, ActivatedRoute } from '@angular/router';

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
  providers: [PostService, UserService],
})
export class DashboardComponent implements OnInit {
  posts: Post[] = [];
  followingCount: number = 0;
  recentFollowing: any[] = [];
  user: string;
  followingMap: { [username: string]: boolean } = {};

  constructor(
    private postsService: PostService,
    private userService: UserService,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    // this.postsService.getPosts().subscribe((data) => {
    //   this.posts = data;
    //   // console.log(this.posts);
    // });

    this.route.paramMap.subscribe((params) => {
      this.user = localStorage.getItem('user') as string;
    });

    this.userService
      .getFollowList(this.user, 'following')
      .subscribe((users) => {
        this.followingCount = users.length;
        this.recentFollowing = users.slice(-4).reverse();
      });

    this.postsService.getDashboardPosts(this.user).subscribe((posts) => {
      this.posts = posts;
      console.log('posts', posts);
    });
  }

  checkIfFollowing(username: string): void {
    this.userService.getFollowList(this.user, 'following').subscribe({
      next: (users) => {
        this.followingMap[username] =
          users?.some((u) => u?.username === username) ?? false;
      },
      error: () => {
        this.followingMap[username] = false;
      },
    });
  }

  toggleFollow(username: string): void {
    this.userService.followUser(this.user, username).subscribe({
      next: () => {
        //this.checkIfFollowing(username);
        this.followingMap[username] = !this.followingMap[username];
      },
      error: (err) => {
        console.error('Error toggling follow for', username, err);
      },
    });
  }
}
