/// <reference types="cypress" />


describe('Lease', () => {
    it("selects property to get a lease", () => {
    cy.visit('http://localhost:3000')
      cy.get('#see-listings').should('be.visible');
      cy.get('#see-listings').click();

      cy.get('#property').should('be.visible');
      cy.get('#property').first().click();
      
      cy.get('#registerInputName').should('be.visible');
      cy.get('#registerInputName').type('Lahari');
      cy.get('#registerInputEmail').type('lbarad@blahblah.com');
      cy.get('#registerInputStartDate').type('2022-03-29');
      cy.get('#registerInputEndDate').type('2023-03-29');
      cy.get('#terms-checkbox').click();
      cy.get('#registerSubmit').click();


      cy.get('#confirmation-msg').should('be.visible');
      cy.get('#print-doc').click();
    });

});