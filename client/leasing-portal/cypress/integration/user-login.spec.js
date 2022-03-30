/// <reference types="cypress" />


describe('Login', () => {
    it("succesfully logins a user", () => {
    cy.visit('http://localhost:3000')
      cy.get('#login').should('be.visible');
      cy.get('#login').click();
  
      cy.get('#login-form').find('[type="email"]').type('testUser@test.com')
      cy.get('#login-form').find('[type="password"]').type('testPasswd');
      cy.get('#loginSubmit').click();
      cy.get('#logout-button').should('be.visible');
    });

    it('Displays User Information', () => {
        cy.get('#my-account').should('be.visible');
        cy.get('#my-account').click();
    })
});