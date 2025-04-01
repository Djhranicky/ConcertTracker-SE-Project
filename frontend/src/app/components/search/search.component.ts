import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { Card } from 'primeng/card';
import { NavbarComponent } from '../navbar/navbar.component';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-search',
  imports: [ProgressSpinnerModule, Card, NavbarComponent, CommonModule],
  templateUrl: './search.component.html',
  styleUrl: './search.component.css',
})
export class SearchComponent implements OnInit {
  query: string = '';
  results: any[] = [];
  loading: boolean = false;
  error: string = '';

  constructor(private route: ActivatedRoute, private http: HttpClient) {}

  ngOnInit() {
    this.route.queryParams.subscribe((params) => {
      this.query = params['q'] || '';
      if (this.query) {
        this.searchArtists();
      }
    });
  }

  searchArtists() {
    this.loading = true;
    const headers = new HttpHeaders().set('Content-Type', 'application/json');
    this.http
      .post(
        `http://localhost:8080/api/artist?name=${this.query}`,
        {}
        // { headers }
      )
      .subscribe({
        next: (data: any) => {
          console.log(data);
          this.results = [data];
          console.log(this.results);
          this.loading = false;
        },
        error: () => {
          this.error = 'No results found.';
          this.loading = false;
        },
      });
  }
}
