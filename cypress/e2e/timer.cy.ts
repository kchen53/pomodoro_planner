describe('Timer', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:8081/session')
    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('H').type('{selectall}{backspace}').type('0')

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('M').type('{selectall}{backspace}').type('0')

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('S').type('{selectall}{backspace}').type('5')

    cy.get('button').contains('Start').click()
    cy.get('div[class="timer-value"]').contains('00:00:05')
    cy.wait(5*1000)

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').should('be.visible')

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('H').type('{selectall}{backspace}').type('0')

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('M').type('{selectall}{backspace}').type('1')

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('S').type('{selectall}{backspace}').type('5')

    cy.get('button').contains('Start').click()
    cy.get('div[class="timer-value"]').contains('00:01:05')
    cy.wait(5*1000)
    cy.get('div[class="timer-value"]').contains('00:00:59')
    cy.get('button').contains('Stop').click()

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').should('be.visible')

    cy.get('button').contains('Reset').click()

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('H').type('{selectall}{backspace}').type('1')

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('M').type('{selectall}{backspace}').type('0')

    cy.get('body').get('app-sessions').get('div[class="body"]').get('div[class="right"]').get('div').get('div[class="timer"]').get('app-timer')
    .get('mat-card').get('mat-card').get('div[class="timer-set ng-star-inserted"]').contains('S').type('{selectall}{backspace}').type('5')

    cy.get('button').contains('Start').click()
    cy.get('div[class="timer-value"]').contains('01:00:05')
    cy.wait(5*1000)
    cy.get('div[class="timer-value"]').contains('00:59:59')
  })
})
