describe('Landing Page', () => {
  it('Visits the initial project page', () => {
    cy.visit('/');
    cy.contains('Concerto');
  });
});
