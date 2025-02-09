import { TestBed } from '@angular/core/testing';

import { PopToursService } from './pop-tours.service';

describe('PopToursService', () => {
  let service: PopToursService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PopToursService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
