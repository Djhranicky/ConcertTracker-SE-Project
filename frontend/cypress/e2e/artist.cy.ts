describe('Artist spec', () => {
  beforeEach(() => {
    cy.visit('/artists');
  });

  it('should display artist details', () => {
    //artist name
    cy.get('.header-artist').should('be.visible');
    //artist img
    cy.get('.header-img').should('have.attr', 'src');
    // shows and tours
    cy.get('.number-box').should('be.visible');
  });

  it('should display recent shows', () => {
    cy.contains('Recent Shows').should('be.visible');
    //show 5 recent
    cy.get('p-carousel').first().find('img').should('have.length', 5);
    //show date
    cy.get('p-tag').first().should('be.visible');
    //show tour name
    //show tour venue
  });

  it('should display upcoming shows', () => {
    cy.contains('Upcoming Shows').should('be.visible');
    //show 3 upcoming
    cy.get('.upcoming p-carousel').find('.stat-frame').should('have.length', 3);
    //show date
    //show tour venue
  });

  it('should display artist stats', () => {
    cy.contains('Stats').should('be.visible');
    cy.contains('Most Played Songs').should('be.visible');
    //show 3 most played songs
    cy.get('p-table').find('tr').should('have.length', 3);
  });
});
