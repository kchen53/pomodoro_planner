describe('about-us page', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:8081/about-us')
    cy.get('p').contains("about-us works!")
  })
})