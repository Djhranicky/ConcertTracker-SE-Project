import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { LandingComponent } from '../landing/landing.component';
import { AuthenticationService } from '../../services/authentication.service';
@Component({
  selector: 'app-home',
  imports: [CommonModule, DashboardComponent, LandingComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  constructor(private authenticationService: AuthenticationService) {}
  isLoggedIn = false;

  ngOnInit() {
    this.isLoggedIn = this.authenticationService.isAuthenticated();
  }
}
