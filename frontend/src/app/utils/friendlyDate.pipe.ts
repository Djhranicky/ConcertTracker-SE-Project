import { Pipe, PipeTransform } from '@angular/core';

@Pipe({ name: 'friendlyDate' })
export class FriendlyDatePipe implements PipeTransform {
  transform(value: string | null): string {
    if (!value) return '';
    const [day, month, year] = value.split('-').map(Number);
    const date = new Date(year, month - 1, day);
    return date.toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    });
  }
}
