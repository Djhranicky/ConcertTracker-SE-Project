import { Component } from '@angular/core';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { FloatLabel } from 'primeng/floatlabel';
import { AuthenticationService } from '../../services/authentication.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    CommonModule,
    InputTextModule,
    ButtonModule,
    FormsModule,
    NavbarComponent,
    FloatLabel,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent {
  email = '';
  password = '';

  constructor(
    private authenticationService: AuthenticationService,
    private router: Router
  ) {}

  login(): void {
    this.authenticationService.login(this.email, this.password).subscribe({
      next: (response) => {
        alert('Login successful!');
        this.router.navigate(['/dashboard']);
      },
      error: (error) => {
        alert('Login failed: ' + error.error.message);
      },
    });
  }
}
