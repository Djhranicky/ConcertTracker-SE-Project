import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-discuss-landing',
  imports: [],
  templateUrl: './discuss-landing.component.html',
  styleUrl: './discuss-landing.component.css',
})
export class DiscussLandingComponent {
  @Input() title: string = '';
  @Input() subtitle: string = '';
  @Input() paragraph: string = '';
}
