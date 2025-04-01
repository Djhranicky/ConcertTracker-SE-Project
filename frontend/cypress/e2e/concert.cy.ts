describe('Concert Page spec', () => {
  beforeEach(() => {
    cy.visit('/concerts');
  });

  it('should display concert header', () => {
    cy.get('.header-artist').should('be.visible');
    cy.get('.header-tour').should('be.visible');
    cy.get('.header-venue').should('be.visible');
    cy.get('.date-frame').should('contain.text', ' ');
  });

  it('should mark concert as attended', () => {
    //test-button
    cy.get('p-button[label="Mark as Attended"]').click();
  });

  it('should display the setlist', () => {
    cy.get('.setlist ol li').should('have.length.greaterThan', 0);
  });

  it('should display recent activity', () => {
    cy.get('.recent p-timeline').should('be.visible');
  });

  it('should display Spotify playlist', () => {
    cy.get('.spotify iframe')
      .should('have.attr', 'src')
      .and('include', 'open.spotify.com');
  });

  it('should display attended users', () => {
    cy.get('.attended p-avatar').should('have.length.greaterThan', 0);
  });

  it('should display other shows and festivals', () => {
    cy.get('.other .p-other').should('contain.text', '2025 Shows');
    cy.get('.other p-avatar-group p-avatar').should(
      'have.length.greaterThan',
      0
    );
  });
});
