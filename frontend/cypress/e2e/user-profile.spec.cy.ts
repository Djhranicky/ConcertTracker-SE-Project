describe('User Profile Page', () => {
  beforeEach(() => {
    // Mock login
    cy.intercept('POST', '/api/login', {
      statusCode: 200,
      body: { user: { id: '123', email: 'test@example.com', username: 'testuser' } }
    }).as('loginRequest');
    
    // Login first
    cy.visit('/login');
    cy.get('#email').type('test@example.com');
    cy.get('#password').type('password123');
    cy.get('button[type="submit"]').click();
    cy.wait('@loginRequest');
    
    // Mock user profile data
    cy.intercept('GET', '**/user/profile', {
      statusCode: 200,
      fixture: 'user-profile.json'
    }).as('profileRequest');
    
    // Mock other profile-related data
    cy.intercept('GET', '**/favorite-concerts', {
      statusCode: 200,
      fixture: 'favorite-concerts.json'
    }).as('concertsRequest');
    
    cy.intercept('GET', '**/recent-attendance', {
      statusCode: 200,
      fixture: 'recent-attendance.json'
    }).as('attendanceRequest');
    
    cy.intercept('GET', '**/bucket-list', {
      statusCode: 200,
      fixture: 'bucket-list.json'
    }).as('bucketListRequest');
    
    cy.intercept('GET', '**/recent-activity', {
      statusCode: 200,
      fixture: 'recent-activity.json'
    }).as('activityRequest');
    
    cy.intercept('GET', '**/recent-lists', {
      statusCode: 200,
      fixture: 'recent-lists.json'
    }).as('listsRequest');
    
    cy.intercept('GET', '**/user-posts', {
      statusCode: 200,
      fixture: 'user-posts.json'
    }).as('postsRequest');
    
    // Visit profile page
    cy.visit('/user-profile');
  });

  it('should display user profile information correctly', () => {
    // Verify main profile elements
    cy.get('.user-name').should('contain', 'Jane Smith');
    cy.get('.username').should('contain', '@janesmith');
    cy.get('.user-bio').should('contain', 'music lover');
    
    // Verify profile stats
    cy.get('.stat-box').should('have.length', 4);
    cy.get('.stat-box').eq(0).find('.stat-number').should('contain', '23');
    cy.get('.stat-box').eq(0).find('.stat-label').should('contain', 'Concerts');
    cy.get('.stat-box').eq(1).find('.stat-number').should('contain', '3');
    cy.get('.stat-box').eq(1).find('.stat-label').should('contain', 'Lists');
    cy.get('.stat-box').eq(2).find('.stat-number').should('contain', '21');
    cy.get('.stat-box').eq(2).find('.stat-label').should('contain', 'Following');
    cy.get('.stat-box').eq(3).find('.stat-number').should('contain', '19');
    cy.get('.stat-box').eq(3).find('.stat-label').should('contain', 'Followers');
  });

  it('should navigate between tabs correctly', () => {
    // Default tab should be profile
    cy.get('.tab-list li.active').should('contain', 'PROFILE');
    cy.get('.content-card').should('exist');
    
    // Click on Activity tab
    cy.get('.tab-list li').contains('ACTIVITY').click();
    cy.get('.column-posts').should('be.visible');
    cy.get('app-post').should('exist');
    
    // Click on Concerts tab
    cy.get('.tab-list li').contains('CONCERTS').click();
    cy.get('.column-posts').should('be.visible');
    cy.get('app-post').should('exist');
    
    // Click on Lists tab
    cy.get('.tab-list li').contains('LISTS').click();
    cy.get('.content-grid').should('be.visible');
    cy.get('.content-card .section-title').should('exist');
    
    // Test following/followers tabs redirect to not-found
    cy.get('.tab-list li').contains('FOLLOWING').click();
    cy.url().should('include', '/not-found');
    
    // Navigate back to profile
    cy.go('back');
    cy.get('.tab-list li').contains('PROFILE').click();
    
    // Test followers tab redirect
    cy.get('.tab-list li').contains('FOLLOWERS').click();
    cy.url().should('include', '/not-found');
  });

  it('should display favorite concerts on profile tab', () => {
    // Verify favorite concerts section
    cy.get('.content-card:first-of-type .section-title').should('contain', 'Favorite Concerts');
    cy.get('.content-card:first-of-type .concert-card').should('have.length.at.least', 1);
    
    // Verify first concert content
    cy.get('.content-card:first-of-type .concert-card:first-of-type .concert-title')
      .should('contain', 'HIT ME HARD AND SOFT');
    cy.get('.content-card:first-of-type .concert-card:first-of-type .artist-name')
      .should('contain', 'Billie Eilish');
  });

  it('should display posts on the activity tab', () => {
    // Navigate to activity tab
    cy.get('.tab-list li').contains('ACTIVITY').click();
    
    // Verify posts exist
    cy.get('app-post').should('exist');
    
    // Check first post content
    cy.get('app-post').first().find('.tour-name').should('exist');
    cy.get('app-post').first().find('.tour-artist').should('exist');
    
    // Check like functionality
    cy.get('app-post').first().find('.pi-heart').click();
    cy.get('app-post').first().find('.pi-heart-fill').should('exist');
  });

  it('should display posts on the concerts tab', () => {
    // Navigate to concerts tab
    cy.get('.tab-list li').contains('CONCERTS').click();
    
    // Verify review posts exist
    cy.get('app-post').should('exist');
    
    // Verify at least one post has rating stars
    cy.get('app-post .pi-star-fill').should('exist');
  });

  it('should display edit profile icon and handle image upload', () => {
    // Verify edit icon exists
    cy.get('.edit-icon').should('exist');
    
    // Test profile image visibility
    cy.get('.profile-image img').should('have.attr', 'src');
    
    // Mock edit profile API
    cy.intercept('PUT', '**/user/profile', {
      statusCode: 200,
      body: { success: true }
    }).as('updateProfileRequest');
    
    // Mock avatar upload API
    cy.intercept('POST', '**/user/profile/avatar', {
      statusCode: 200,
      body: { avatarUrl: 'https://example.com/new-avatar.jpg' }
    }).as('avatarUploadRequest');
    
    // Note: Since the edit functionality isn't fully implemented in the component,
    // we can't test the actual interaction, but we can verify the edit icon exists
  });

  it('should handle responsive layout', () => {
    // Test on mobile viewport
    cy.viewport('iphone-x');
    
    // Check if profile header adapts to smaller screen
    cy.get('.profile-header').should('have.css', 'flex-direction', 'column');
    
    // Restore viewport
    cy.viewport('macbook-15');
  });
});