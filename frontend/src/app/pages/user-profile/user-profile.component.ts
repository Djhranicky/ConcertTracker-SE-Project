import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavbarComponent } from '../../components/navbar/navbar.component';

@Component({
  selector: 'app-user-profile',
  standalone: true,
  imports: [CommonModule, FormsModule, NavbarComponent],
  templateUrl: './user-profile.component.html',
  styleUrl: './user-profile.component.css'
})
export class UserProfileComponent implements OnInit {
  // Active tab state
  activeTab: string = 'profile';

  // Mock data for user profile
  userData = {
    name: 'Jane Smith',
    bio: '24. music lover. user description.',
    profileImage: 'assets/images/user-profile.jpg',
    stats: {
      concerts: 23,
      lists: 3,
      following: 21,
      followers: 19
    }
  };

  // Mock data for favorite concerts
  favoriteConcerts = [
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    }
  ];

  // Mock data for recent attendance
  recentAttendance = [
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    }
  ];

  // Mock data for bucket list
  bucketList = [
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    },
    {
      title: 'HIT ME HARD AND SOFT',
      artist: 'Billie Eilish',
      date: 'Feb 19, 2025',
      image: 'assets/images/billie-eilish.jpg'
    }
  ];

  // Mock data for recent activity with HTML formatting
  recentActivity = [
    { 
      text: 'You followed <span class="highlight">John Doe</span>'
    },
    { 
      text: 'You added a show to <span class="highlight">2025 shows</span> list'
    },
    { 
      text: 'You attended Billie Eilish\'s <span class="highlight">HIT ME HARD AND SOFT</span>'
    }
  ];

  // Mock data for recent lists
  recentLists = {
    shows: [
      'assets/images/post-malone.jpg',
      'assets/images/concert-poster.jpg'
    ],
    festivals: [
      'assets/images/post-malone.jpg',
      'assets/images/concert-poster.jpg'
    ]
  };

  constructor() { }

  ngOnInit(): void {
    // Any initialization logic goes here
  }

  // Method to change active tab
  setActiveTab(tab: string): void {
    this.activeTab = tab;
  }
}