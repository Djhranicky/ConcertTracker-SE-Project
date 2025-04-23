import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { PostComponent } from '../../components/post/post.component';
import { UserService } from '../../services/user.service';
import {
  UserProfile,
  ConcertCard,
  Activity,
  List,
} from '../../models/user.model';
import { Post } from '../../models/post.model';
import { Card } from 'primeng/card';
import { AvatarModule } from 'primeng/avatar';
import { Router } from '@angular/router';
import { ProgressSpinner } from 'primeng/progressspinner';
import { RouterModule, ActivatedRoute } from '@angular/router';
import { Button } from 'primeng/button';

@Component({
  selector: 'app-user-profile',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    NavbarComponent,
    PostComponent,
    Card,
    AvatarModule,
    ProgressSpinner,
    RouterModule,
    Button,
  ],
  templateUrl: './user-profile.component.html',
  styleUrl: './user-profile.component.css',
  providers: [UserService],
})
export class UserProfileComponent implements OnInit {
  user: string;
  isUser: boolean = false;
  loading: boolean = false;
  loggedInUser: string;
  isFollowing: boolean = false;
  followingCount: number;
  followersCount: number;
  // Active tab state
  activeTab: string = 'profile';

  // User profile data
  userProfile!: UserProfile;
  favoriteConcerts: ConcertCard[] = [];
  recentAttendance: ConcertCard[] = [];
  bucketList: ConcertCard[] = [];
  recentActivity: Activity[] = [];
  recentLists: List[] = [];

  // Posts for concert-related tabs
  userPosts: Post[] = [];
  favoritePosts: Post[] = [];
  bucketListPosts: Post[] = [];

  constructor(
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      this.user = params.get('user') as string;
      this.loggedInUser = localStorage.getItem('user') as string;
      this.isUser = this.user === this.loggedInUser;
      console.log(this.user, this.loggedInUser);
      this.loading = true;

      this.checkIfFollowing();
    });

    this.userService
      .getFollowList(this.user, 'following')
      .subscribe((users) => {
        this.followingCount = users.length;
      });

    this.userService
      .getFollowList(this.user, 'followers')
      .subscribe((users) => {
        this.followersCount = users.length;
      });

    // Get user profile data
    this.userService.getUserProfile().subscribe((profile) => {
      this.userProfile = profile;
    });

    // Get concert data
    this.userService.getFavoriteConcerts().subscribe((concerts) => {
      this.favoriteConcerts = concerts;
    });

    this.userService.getRecentAttendance().subscribe((concerts) => {
      this.recentAttendance = concerts;
    });

    this.userService.getBucketList().subscribe((concerts) => {
      this.bucketList = concerts;
    });

    this.userService.getRecentActivity().subscribe((activities) => {
      this.recentActivity = activities;
    });

    this.userService.getRecentLists().subscribe((lists) => {
      this.recentLists = lists;
    });

    // Get posts for concert display
    this.userService.getUserPosts().subscribe((posts) => {
      this.userPosts = posts;

      // Filter posts by type for different tabs
      this.favoritePosts = posts.filter((post) => post.type === 'review');
      this.bucketListPosts = posts.filter((post) => post.type === 'wishlist');
    });
  }

  // Method to change active tab
  setActiveTab(tab: string): void {
    if (tab === 'following' || tab === 'followers') {
      // Navigate to Not Found page instead of displaying the tab
      this.router.navigate(['/not-found']);
    } else {
      this.activeTab = tab;
    }
  }

  toggleFollow() {
    console.log('logged in', this.loggedInUser);
    this.userService.followUser(this.loggedInUser, this.user).subscribe({
      next: () => {
        console.log('followed!');
        this.checkIfFollowing();
      },
      error: (err) => {
        console.error('Error following user', err);
      },
    });
  }

  checkIfFollowing(): void {
    this.userService
      // .getFollowList(this.loggedInID, 'following')
      .getFollowList(this.loggedInUser, 'following')
      .subscribe({
        next: (users) => {
          if (!users) {
            this.isFollowing = false;
            return;
          }
          this.isFollowing = users.some((u) => u?.username === this.user);
        },
        error: (err) => {
          this.isFollowing = false;
        },
      });
  }
}
