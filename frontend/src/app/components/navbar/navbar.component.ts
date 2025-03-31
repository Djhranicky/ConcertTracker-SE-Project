import { Component, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { Menubar } from 'primeng/menubar';
import { Menu } from 'primeng/menu';
import { BadgeModule } from 'primeng/badge';
import { AvatarModule } from 'primeng/avatar';
import { InputTextModule } from 'primeng/inputtext';
import { CommonModule } from '@angular/common';
import { ButtonModule } from 'primeng/button';
import { RouterModule } from '@angular/router';
import { AuthenticationService } from '../../services/authentication.service';
import { Ripple } from 'primeng/ripple';

@Component({
  selector: 'app-navbar',
  imports: [
    Menu,
    Menubar,
    BadgeModule,
    AvatarModule,
    InputTextModule,
    CommonModule,
    ButtonModule,
    RouterModule,
    Ripple,
  ],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css',
})
export class NavbarComponent {
  constructor(
    private authenticationService: AuthenticationService,
    private router: RouterModule
  ) {}
  isLoggedIn: boolean = false;
  items: MenuItem[] | undefined;
  userItems: MenuItem[] | undefined;
  logIn = {
    colorScheme: {
      light: {
        root: {
          primary: {
            color: 'black',
          },
        },
      },
    },
  };

  logout() {
    this.authenticationService.logout();
  }

  ngOnInit() {
    this.authenticationService.isAuthenticated().subscribe((isAuth) => {
      this.isLoggedIn = isAuth;
      if (this.isLoggedIn) {
        this.items = [
          {
            label: 'Home',
            routerLink: '/',
          },
          {
            label: 'Concerts',
            routerLink: '/concerts',
          },
          {
            label: 'Artists',
            routerLink: '/artists',
          },
          {
            label: 'Lists',
            routerLink: '/lists',
          },
        ];

        this.userItems = [
          {
            label: 'Profile',
            routerLink: '/user-profile',
          },
          {
            label: 'Notifications',
            routerLink: '/notifications',
          },
          {
            label: 'Settings',
            routerLink: '/settings',
          },
          {
            label: 'Sign out',
            command: (event) => {
              this.logout();
            },
          },
        ];
      }
    });
  }
}
