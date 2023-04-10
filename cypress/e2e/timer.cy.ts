describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:8081/session')
    cy.get('input[placeholder="mm:ss"]').type('00:20')
    cy.get('button').contains('Start').click()
    cy.get('div[class="timer-value"]').contains('00:20')
    cy.wait(20*1000)
    cy.get('div[class="timer-value"]').contains('00:00')
  })
})
