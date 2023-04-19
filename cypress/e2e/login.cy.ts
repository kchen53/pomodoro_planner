describe('login page', () => {
  beforeEach(() => {
    cy.visit('http://localhost:8081/home')
  })
  it('passes', () => {
    cy.contains('Login').click()
    cy.contains('Sign Up').click()
    cy.get('input[name="username"]').type('Admin')
    cy.get('input[name="password"]').type('test')
    cy.get('button').contains('Submit').click()
  })
})
