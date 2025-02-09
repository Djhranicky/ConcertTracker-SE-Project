import { Component } from '@angular/core';
import { LandingHeroComponent } from '../../components/landing-hero/landing-hero.component';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { PopCarouselComponent } from '../../components/pop-carousel/pop-carousel.component';
import { FeaturesLandingComponent } from '../../components/features-landing/features-landing.component';
import { DiscussLandingComponent } from '../../components/discuss-landing/discuss-landing.component';
import { DiscoverLandingComponent } from '../../components/discover-landing/discover-landing.component';
import { TweetsLandingComponent } from '../../components/tweets-landing/tweets-landing.component';
@Component({
  selector: 'app-landing',
  imports: [
    LandingHeroComponent,
    NavbarComponent,
    PopCarouselComponent,
    FeaturesLandingComponent,
    DiscussLandingComponent,
    DiscoverLandingComponent,
    TweetsLandingComponent,
  ],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.css',
})
export class LandingComponent {}
