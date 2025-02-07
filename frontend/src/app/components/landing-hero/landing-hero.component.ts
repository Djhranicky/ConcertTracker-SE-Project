import { Component } from '@angular/core';
import { ButtonModule } from 'primeng/button';
import { RippleModule } from 'primeng/ripple';

@Component({
  selector: 'app-landing-hero',
  imports: [ButtonModule, RippleModule],
  templateUrl: './landing-hero.component.html',
  styleUrl: './landing-hero.component.css',
})
export class LandingHeroComponent {}
