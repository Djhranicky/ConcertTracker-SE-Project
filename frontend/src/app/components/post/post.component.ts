import { Component, Input } from '@angular/core';
import { PostService, Post } from '../../services/post.service';
import { Button } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { Avatar } from 'primeng/avatar';
@Component({
  selector: 'app-post',
  imports: [CardModule, Button, Avatar],
  templateUrl: './post.component.html',
  styleUrl: './post.component.css',
  providers: [PostService],
})
export class PostComponent {
  @Input() post: Post;
  // responsiveOptions: any[] | undefined;
}
