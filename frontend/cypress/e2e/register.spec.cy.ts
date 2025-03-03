describe('Register Page', () => {
  beforeEach(() => {
    cy.visit('/register');
  });
  it('Visits the Register page', () => {
    cy.visit('/register');
    cy.contains('Sign up');
  });

  it('should render username and password input fields and login button', () => {
    cy.get('#email').should('exist');
    cy.get('#username').should('exist');
    cy.get('#password').should('exist');
    cy.get('#confirmPassword').should('exist');
    cy.get('button[type="submit"]').should('exist');
  });

  it('should have email, user and passwords as required fields of form', () => {
    cy.get('#email').click();
    cy.get('#username').click();
    cy.get('#password').click();
    cy.get('#confirmPassword').type(' ');
    cy.get('#email').click();

    cy.contains('Please enter a valid email.').should('be.visible');
    cy.contains('Please enter a username').should('be.visible');
    cy.contains('Please enter password with at least 6 characters').should(
      'be.visible'
    );
    cy.contains('Please enter matching password').should('be.visible');
  });

  it('should show validation error for invalid email format', () => {
    cy.get('#email').type('invalid');
    cy.get('#password').click();

    cy.contains('Please enter a valid email.').should('be.visible');
  });

  it('should submit form and navigate to login page on successful registration', () => {
    cy.intercept('POST', '/api/register', {
      statusCode: 200,
    }).as('registerRequest');

    cy.get('#email').type('test@example.com');
    cy.get('#username').type('test');
    cy.get('#password').type('password123');
    cy.get('#confirmPassword').type('password123');

    cy.get('button[type="submit"]').click();
    cy.wait('@registerRequest');

    cy.url().should('include', '/login');
  });

  //should show error if registration failed

  it('should redirect to /login when clicking log in here on description', () => {
    cy.get('a.url').click();
    cy.url().should('include', '/login');
  });
});
