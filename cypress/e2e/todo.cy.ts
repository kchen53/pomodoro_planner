describe('todo page', () => {
  beforeEach(() => {
    cy.visit('http://127.0.0.1:8081/session')
  })
  it('passes', () => {
    cy.get('input[placeholder="task..."]').type('Check this box')
    cy.get('button[name="Add"]').click()
    cy.get('[type="checkbox"]').check()
    cy.get('button[name="Delete"]').last().click()
  })
})
