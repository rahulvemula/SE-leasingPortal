/// <reference types="cypress" />


describe('Login', () => {
    it("succesfully logins a user", () => {
    cy.visit('http://localhost:3000/SE-leasingPortal')
      cy.get('#login').should('be.visible');
      cy.get('#login').click();
  
      cy.get('#login-form').find('[type="email"]').type('testUser@test.com')
      cy.get('#login-form').find('[type="password"]').type('testPasswd');
      cy.get('#loginSubmit').click();
      cy.get('#logout-button').should('be.visible');
    });

    it('Displays My Account page', () => {
        cy.get('#my-account').should('be.visible');
        cy.get('#my-account').click();        
    });

    it('Displays User information', () => {
      cy.get('#info-name').should('be.visible');
      cy.get('#info-name').should('have.text', 'Test User');

      cy.get('#info-email').should('be.visible');
      cy.get('#info-email').should('have.text', 'testUser@test.com');

    })
});