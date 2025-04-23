import { Component } from '@angular/core';
import { SearchComponent } from '../../components/search/search.component';
import { NavbarComponent } from '../../components/navbar/navbar.component';
@Component({
  selector: 'app-search-page',
  imports: [SearchComponent, NavbarComponent],
  templateUrl: './search.component.html',
  styleUrl: './search.component.css',
})
export class SearchPage {}
