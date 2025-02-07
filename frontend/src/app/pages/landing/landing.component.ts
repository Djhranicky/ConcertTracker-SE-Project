import { Component } from '@angular/core';
import { LandingHeroComponent } from '../../components/landing-hero/landing-hero.component';
import { NavbarComponent } from '../../components/navbar/navbar.component';
@Component({
  selector: 'app-landing',
  imports: [LandingHeroComponent, NavbarComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.css',
})
export class LandingComponent {}
