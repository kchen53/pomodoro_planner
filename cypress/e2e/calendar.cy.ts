describe('Calendar', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:8081/')

    cy.get('app-calendar').should('be.visible')

  })
})
