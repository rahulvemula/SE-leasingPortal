/// <reference types="cypress" />

describe('Shows terms and conditions', () => {
    it("renders terms page from listing page", () => {
        cy.visit('http://localhost:3000/SE-leasingPortal/listing/1');

        cy.get('#terms').should('be.visible');
        cy.get('#terms').click();

        cy.get('#tandc').should('be.visible');
    })
});