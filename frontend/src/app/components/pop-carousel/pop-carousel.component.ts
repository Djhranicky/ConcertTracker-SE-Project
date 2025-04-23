import { Component, OnInit } from '@angular/core';
import { PopToursService } from '../../services/pop-tours.service';
import { Tour } from '../../services/concert.service';
import { CarouselModule } from 'primeng/carousel';
import { ButtonModule } from 'primeng/button';
import { TagModule } from 'primeng/tag';

@Component({
  selector: 'app-pop-carousel',
  imports: [CarouselModule, ButtonModule, TagModule],
  templateUrl: './pop-carousel.component.html',
  styleUrl: './pop-carousel.component.css',
  providers: [PopToursService],
})
export class PopCarouselComponent implements OnInit {
  tours: Tour[] = [];

  responsiveOptions: any[] | undefined;

  constructor(private popTourService: PopToursService) {}

  ngOnInit() {
    this.popTourService
      .getPopTours()
      .toPromise()
      .then((tours) => {
        this.tours = tours!;
      });

    this.responsiveOptions = [
      {
        breakpoint: '1400px',
        numVisible: 5,
        numScroll: 1,
      },
      {
        breakpoint: '1199px',
        numVisible: 3,
        numScroll: 1,
      },
      {
        breakpoint: '850px',
        numVisible: 2,
        numScroll: 1,
      },
      {
        breakpoint: '575px',
        numVisible: 1,
        numScroll: 1,
      },
    ];
  }
}
