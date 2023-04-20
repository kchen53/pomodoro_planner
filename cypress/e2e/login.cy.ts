describe('Login', () => {
  beforeEach(() => {
    cy.visit('http://localhost:8081/home')
  })
  it('passes', () => {
    cy.contains('Login').click()
    cy.get('input[name="username"]').type('Username')
    cy.get('input[name="password"]').type('password')
    cy.get('button').contains('Submit').click()

    cy.contains('Sign Up').click()
    cy.get('input[name="username"]').type('Username')
    cy.get('input[name="password"]').type('password')
    cy.get('button').contains('Submit').click()
  })
})
