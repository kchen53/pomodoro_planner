import { LoginComponent } from "./login.component";

describe("LoginComponent", () => {
  it ('mounts', () => {
    cy.mount(LoginComponent)
  })
});
