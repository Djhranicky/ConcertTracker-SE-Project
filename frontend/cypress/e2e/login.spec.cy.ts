describe('Login Page', () => {
  beforeEach(() => {
    cy.visit('/login');
  });
  it('Visits the login page', () => {
    cy.visit('/login');
    cy.contains('Sign in');
  });

  it('should render username and password input fields and login button', () => {
    cy.get('#email').should('exist');
    cy.get('#password').should('exist');
    cy.get('button[type="submit"]').should('exist');
  });

  it('should have email and passwords as required fields of form', () => {
    cy.get('#email').click();
    cy.get('#password').click();
    cy.get('#email').click();

    cy.contains('Please enter a valid email.').should('be.visible');
    cy.contains('Please enter password').should('be.visible');
  });

  it('should show validation error for invalid email format', () => {
    cy.get('#email').type('invalid');
    cy.get('#password').click();

    cy.contains('Please enter a valid email.').should('be.visible');
  });

  it('should submit form and navigate to dashboard on successful login', () => {
    cy.intercept('POST', '/api/login', {
      statusCode: 200,
    }).as('loginRequest');

    cy.get('#email').type('test@example.com');
    cy.get('#password').type('password123');

    cy.get('button[type="submit"]').click();
    cy.wait('@loginRequest');

    cy.url().should('include', '/');
  });

  //should show error if login failed

  it('should redirect to /register when clicking register on description', () => {
    cy.get('a.url').click();
    cy.url().should('include', '/register');
  });
});
