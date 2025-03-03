import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-root',
  imports: [RouterModule],
  providers: [HttpClient],
  template: `<router-outlet></router-outlet>`,
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'concerto';
}
