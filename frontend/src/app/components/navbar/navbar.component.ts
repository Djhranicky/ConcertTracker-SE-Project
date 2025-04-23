import { Component, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { Menubar } from 'primeng/menubar';
import { Menu } from 'primeng/menu';
import { BadgeModule } from 'primeng/badge';
import { AvatarModule } from 'primeng/avatar';
import { InputTextModule } from 'primeng/inputtext';
import { CommonModule } from '@angular/common';
import { ButtonModule } from 'primeng/button';
import { RouterModule, Router } from '@angular/router';
import { AuthenticationService } from '../../services/authentication.service';
import { Ripple } from 'primeng/ripple';
import { FormsModule } from '@angular/forms';
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
    FormsModule,
  ],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css',
})
export class NavbarComponent {
  constructor(
    private authenticationService: AuthenticationService,
    private router: Router
  ) {}
  isLoggedIn: boolean = false;
  items: MenuItem[] | undefined;
  userItems: MenuItem[] | undefined;
  query: string = '';

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

  onSearch() {
    if (this.query.trim()) {
      this.router.navigate(['/search'], { queryParams: { q: this.query } });
    }
  }

  ngOnInit() {
    this.authenticationService.isAuthenticated().subscribe((isAuth) => {
      this.isLoggedIn = isAuth;
      let user = localStorage.getItem('user') as string;
      if (this.isLoggedIn) {
        this.items = [
          {
            label: 'Home',
            routerLink: '/',
          },
          // {
          //   label: 'Concerts',
          //   routerLink: '/concerts',
          // },
          // {
          //   label: 'Artists',
          //   routerLink: '/artists',
          // },
          // {
          //   label: 'Lists',
          //   routerLink: '/lists',
          // },
        ];

        this.userItems = [
          {
            label: 'Profile',
            routerLink: `/user/${user}`,
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
