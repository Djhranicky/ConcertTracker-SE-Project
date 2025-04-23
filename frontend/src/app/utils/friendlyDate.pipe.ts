import { Pipe, PipeTransform } from '@angular/core';

@Pipe({ name: 'friendlyDate' })
export class FriendlyDatePipe implements PipeTransform {
  transform(value: string | null): string {
    if (!value) return '';

    let date: Date;
    if (/^\d{2}-\d{2}-\d{4}$/.test(value)) {
      const [day, month, year] = value.split('-').map(Number);
      date = new Date(year, month - 1, day);
    } else {
      date = new Date(value);
    }
    return date.toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    });
  }
}
