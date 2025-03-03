import { Component, OnInit } from '@angular/core';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { FloatLabel } from 'primeng/floatlabel';
import { MessageModule } from 'primeng/message';
import {
  FormsModule,
  ReactiveFormsModule,
  FormBuilder,
  Validators,
  FormGroup,
} from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { AuthenticationService } from '../../services/authentication.service';
import { Router } from '@angular/router';
import { matchValidator } from '../../utils/match-validator';

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [
    CommonModule,
    InputTextModule,
    ButtonModule,
    FormsModule,
    NavbarComponent,
    FloatLabel,
    ReactiveFormsModule,
    MessageModule,
  ],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css',
})
export class SignupComponent implements OnInit {
  signupForm: FormGroup;

  ngOnInit(): void {
    this.signupForm = this.formBuilder.group(
      {
        email: ['', [Validators.required, Validators.email]],
        username: ['', [Validators.required]],
        password: ['', [Validators.required, Validators.minLength(6)]],
        confirmPassword: ['', [Validators.required]],
      },
      {
        validators: matchValidator('password', 'confirmPassword'),
      }
    );
  }

  constructor(
    private authenticationService: AuthenticationService,
    private router: Router,
    private formBuilder: FormBuilder
  ) {}

  register(): void {
    if (this.signupForm.valid) {
      this.authenticationService
        .register(
          this.signupForm.value.email,
          this.signupForm.value.username,
          this.signupForm.value.password
        )
        .subscribe({
          next: (response) => {
            alert('Registration successful! Please login.');
            this.router.navigate(['/login']);
          },
          error: (error) => {
            alert('Registration failed: ' + error.message);
          },
        });
    }
  }
}
