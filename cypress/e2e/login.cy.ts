describe('about-us page', () => {
  beforeEach(() => {
    cy.visit('http://127.0.0.1:8081/login')
  })
  it('passes', () => {
    cy.get('#login-content').contains('Login');
    cy.get('input[name="username"]').type('user1')
    cy.get('input[name="password"]').type('password')
    cy.get('button[type="submit"]').click()
  })
})
