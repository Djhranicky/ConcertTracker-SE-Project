import { Component } from '@angular/core';
import { ButtonModule } from 'primeng/button';
import { RippleModule } from 'primeng/ripple';
import { StyleClassModule } from 'primeng/styleclass';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-landing-hero',
  imports: [ButtonModule, RippleModule, StyleClassModule, RouterModule],
  templateUrl: './landing-hero.component.html',
  styleUrl: './landing-hero.component.css',
})
export class LandingHeroComponent {}
