import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { PostComponent } from '../../components/post/post.component';
import { UserService, UserProfile, ConcertCard, Activity, List, Following, Followers } from '../../services/user.service';
import { Post } from '../../services/post.service';
// import { Button } from 'primeng/button';
import { Card } from 'primeng/card';
import { AvatarModule } from 'primeng/avatar';
import { Router } from '@angular/router'; 

@Component({
  selector: 'app-user-profile',
  standalone: true,
  imports: [
    CommonModule, 
    FormsModule, 
    NavbarComponent, 
    PostComponent,
    Card,
    AvatarModule
  ],
  templateUrl: './user-profile.component.html',
  styleUrl: './user-profile.component.css',
  providers: [UserService]
})
export class UserProfileComponent implements OnInit {
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

  // Following and followers
  followings: Following[] = [];
  followers: Followers[] = [];

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    // Get user profile data
    this.userService.getUserProfile().subscribe(profile => {
      this.userProfile = profile;
    });

    // Get concert data
    this.userService.getFavoriteConcerts().subscribe(concerts => {
      this.favoriteConcerts = concerts;
    });

    this.userService.getRecentAttendance().subscribe(concerts => {
      this.recentAttendance = concerts;
    });

    this.userService.getBucketList().subscribe(concerts => {
      this.bucketList = concerts;
    });

    this.userService.getRecentActivity().subscribe(activities => {
      this.recentActivity = activities;
    });

    this.userService.getRecentLists().subscribe(lists => {
      this.recentLists = lists;
    });

    this.userService.getFollowers().subscribe(followers => {
      this.followers = followers;
    });

    this.userService.getFollowing().subscribe(following => {
      this.followings = following;
    })

    // Get posts for concert display
    this.userService.getUserPosts().subscribe(posts => {
      this.userPosts = posts;
      
      // Filter posts by type for different tabs
      this.favoritePosts = posts.filter(post => post.type === 'review');
      this.bucketListPosts = posts.filter(post => post.type === 'wishlist');
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
}
