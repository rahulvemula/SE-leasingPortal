/// <reference types="cypress" />


describe('Support', () => {
    it("registers a complaint", () => {
        cy.visit('http://localhost:3000/SE-leasingPortal/support');

        cy.get('#registerInputName').should('be.visible');
        cy.get('#registerInputName').type('Lahari');

        cy.get('#registerInputProperty').should('be.visible');
        cy.get('#registerInputProperty').type('The Niche Student Appartments');

        cy.get('#locations').should('be.visible');
        cy.get('#locations').select('Kitchen');

        cy.get('#priority').should('be.visible');
        cy.get('#priority').select('Medium');

        
        cy.get('#registerInputDescription').should('be.visible');
        cy.get('#registerInputDescription').type('The dishwasher does not work');

        const stub = cy.stub()  
        cy.on ('window:alert', stub)

        cy.get('#registerSubmit').click().then(() => {
            expect(stub.getCall(0)).to.be.calledWith('Complaint registered, we will get back to you!')      
          })  

    });
});