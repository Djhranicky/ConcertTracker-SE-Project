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
  loggedInID: number;
  userID: number = 4;
  isFollowing: boolean = false;
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
      const loggedInUser = localStorage.getItem('user');
      this.loggedInID = Number(localStorage.getItem('id') as string);
      this.isUser = this.user === loggedInUser;
      console.log(this.user, loggedInUser);
      this.loading = true;

      this.checkIfFollowing();
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
    this.userService.followUser(this.loggedInID, this.userID).subscribe({
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
      .getFollowList(this.loggedInID, 'following')
      .subscribe((users) => {
        console.log('users', users);
        this.isFollowing = users.some((u) => u.userName === this.user);
      });
  }
}
