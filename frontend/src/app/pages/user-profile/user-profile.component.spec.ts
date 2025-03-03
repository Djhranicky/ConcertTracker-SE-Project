import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { UserProfileComponent } from './user-profile.component';

describe('UserProfileComponent', () => {
  let component: UserProfileComponent;
  let fixture: ComponentFixture<UserProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserProfileComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display the correct user name', () => {
    const nameElement = fixture.debugElement.query(By.css('.user-name')).nativeElement;
    expect(nameElement.textContent).toContain('Jane Smith');
  });

  it('should have "profile" as the default active tab', () => {
    expect(component.activeTab).toBe('profile');
    const activeTab = fixture.debugElement.query(By.css('.tab-list li.active')).nativeElement;
    expect(activeTab.textContent).toBe('PROFILE');
  });

  it('should change active tab when a tab is clicked', () => {
    const tabs = fixture.debugElement.queryAll(By.css('.tab-list li'));
    const activityTab = tabs[1].nativeElement;
    
    activityTab.click();
    fixture.detectChanges();
    
    expect(component.activeTab).toBe('activity');
    expect(activityTab.classList).toContain('active');
  });

  it('should display the correct number of favorite concerts', () => {
    const concertCards = fixture.debugElement.queryAll(By.css('.content-card:first-of-type .concert-card'));
    expect(concertCards.length).toBe(component.favoriteConcerts.length);
  });

  it('should display the correct number of recent attendance concerts', () => {
    const attendanceCards = fixture.debugElement.queryAll(
      By.css('.content-card:not(:first-of-type) .concert-grid:first-of-type .concert-card')
    );
    expect(attendanceCards.length).toBe(component.recentAttendance.length);
  });

  it('should display the correct statistics text', () => {
    const statsText = fixture.debugElement.queryAll(By.css('.stats-text'));
    expect(statsText[0].nativeElement.textContent).toContain('pop');
    expect(statsText[0].nativeElement.textContent).toContain('indie');
    expect(statsText[1].nativeElement.textContent).toContain('arenas');
    expect(statsText[1].nativeElement.textContent).toContain('17 songs');
  });

  it('should display the correct number of recent activities', () => {
    const activityItems = fixture.debugElement.queryAll(By.css('.activity-item'));
    expect(activityItems.length).toBe(component.recentActivity.length);
  });

  it('should display the bucket list section', () => {
    const bucketListTitle = fixture.debugElement.query(
      By.css('.content-card:last-of-type .section-title')
    ).nativeElement;
    expect(bucketListTitle.textContent).toBe('Bucket List');
    
    const bucketListItems = fixture.debugElement.queryAll(
      By.css('.content-card:last-of-type .concert-card')
    );
    expect(bucketListItems.length).toBe(component.bucketList.length);
  });
});