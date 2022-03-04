/// <reference types="cypress" />


describe('Register', () => {
    beforeEach(() => {
      cy.visit('http://localhost:3000')
    })
  
    it("registers a user", () => {
      cy.get('#register').should('be.visible');
      cy.get('#register').click();
  
      cy.get('#register-form').find('[type="text"]').type('Test User')
      cy.get('#register-form').find('[type="email"]').type('testUser@test.com')
      cy.get('#register-form').find('[type="password"]').type('testPasswd');
      cy.get('#registerSubmit').click();
    });
});