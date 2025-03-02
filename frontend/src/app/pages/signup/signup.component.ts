import { Component } from '@angular/core';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { AuthenticationService } from '../../services/authentication.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [
    CommonModule,
    InputTextModule,
    ButtonModule,
    FormsModule,
    NavbarComponent,
  ],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css',
})
export class SignupComponent {
  email = '';
  username = '';
  password = '';

  constructor(
    private authenticationService: AuthenticationService,
    private router: Router
  ) {}

  register(): void {
    this.authenticationService
      .register(this.email, this.username, this.password)
      .subscribe({
        next: (response) => {
          alert('Registration successful! Please login.');
          this.router.navigate(['/login']);
        },
        error: (error) => {
          console.log(this.email, this.username, this.password);
          alert('Registration failed: ' + error.error.message);
        },
      });
  }
}
