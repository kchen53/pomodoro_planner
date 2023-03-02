# Sprint 2

## Work Completed
- Set up backend server with angular frontend integration
- Made a navbar with mutiple routes
- Created timer, about-us, login, and to-do components
- Made a basic login page that can take input
- Made a basic to-do list that has front-end and back-end qualities (can post and delete)
- Completed To Do list implementation
- Succesfully used mysql to store data
- Wrote API for routes completed so far

## Unit Tests

### Frontend:

#### Unit tests for Angular using Karma and Jasmine:

it('check if about-us page text displays', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('p').textContent).toContain('about-us works!');
  });
  
it('check username and password if they start off empty', () => {
  expect(component.username).toBe("");
  expect(component.password).toBe("");
});

it('login console to be called', () => {
  expect(component.login).toHaveBeenCalled;
});

it('check if timer page text displays', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('p').textContent).toContain('timer works!');
  });
  
it('check if addTodo function was called', () => {
  expect(component.addTodo).toHaveBeenCalled;
});

it('check if getTodo function was called', () => {
  expect(component.getTodos).toHaveBeenCalled;
});

it('check if deleteTodo function was called', () => {
  expect(component.deleteTodo).toHaveBeenCalled;
});

it('check if the to-do list html header', () => {
  fixture.detectChanges();
  const compiled = fixture.nativeElement;
  expect(compiled.querySelector('mat-card-title').textContent).toContain('To-Do List');
});

#### Cypress:

describe('login page', () => {
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

### Backend:

(pomodoro_planner\src\server\pkg\controllers\todo-controller_test.go)
- TestGetToDo
- TestGetToDoByID
- TestCreateToDo
- TestDeleteToDo
- TestUpdateToDo

## API

Pomodoro Planner aids the creation of a study session environment. We make use of a RESTful API to facilitate http requests and responses. Data is accepted and returned in the JSON format.

> Base URL: <br>
> 127.0.0.1:8081/

### Contents  

- [**Todo**](#todo)
    - [Create](#create)  <br>
    - [Get All](#get-all)  <br>
    - [Get By ID](#get-by-id)  <br>
    - [Update](#update)  <br>
    - [Delete](#delete) <br>
- [**Login**](#login)

---

### Todo

####  Create

> <font color="orange">POST</font> 127.0.0.1:8081/todo <br>

Creates a new Todo item, returns the created Todo item<br>

##### Input Fields:

>> **"id"** number <br>
>> **"task"** string <br>
>> **"due"** string <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Input:

```json
{
    "id":1,
    "task":"Pay pomos",
    "due":"Friday",
    "complete":false
}
```
 
##### Output Fields:

>> **"ID"** number :warning:Not currently used <br>
>> **"CreatedAt"** string :warning:Not currently used <br>
>> **"UpdatedAt"** string :warning:Not currently used <br>
>> **"DeletedAt"** string :warning:Not currently used <br>
>> **"id"** number <br>
>> **"task"** string <br>
>> **"due"** string <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
{
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "id": 1,
    "task": "Pay pomos",
    "due": "Friday",
    "complete": false
}
```
          
#### Get All

> GET 127.0.0.1:8081/todo <br>

Returns a list of all stored Todo items <br>

##### Output Fields:

>> **"ID"** number :warning:Not currently used <br>
>> **"CreatedAt"** string :warning:Not currently used <br>
>> **"UpdatedAt"** string :warning:Not currently used <br>
>> **"DeletedAt"** string :warning:Not currently used <br>
>> **"id"** number <br>
>> **"task"** string <br>
>> **"due"** string <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
[
    {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 1,
        "task": "Pay pomos",
        "due": "Friday",
        "complete": false
    },
    {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 2,
        "task": "Release doros",
        "due": "Tuesday",
        "complete": false
    }
]
```

#### Get by ID

> GET 127.0.0.1:8081/todo/{id} <br>

Returns the Todo item with the id designated by the address

##### Output Fields:

>> **"ID"** number :warning:Not currently used <br>
>> **"CreatedAt"** string :warning:Not currently used <br>
>> **"UpdatedAt"** string :warning:Not currently used <br>
>> **"DeletedAt"** string :warning:Not currently used <br>
>> **"id"** number <br>
>> **"task"** string <br>
>> **"due"** string <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

> GET 127.0.0.1:8081/todo/1

```json
{
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "id": 1,
    "task": "Pay pomos",
    "due": "Friday",
    "complete": false
}
```

#### Update

> PUT 127.0.0.1:8081/todo/{id} <br>

Modifies the Todo item with the id designated by the address, returns the created Todo item<br>

##### Input Fields:

>> **"id"** number <br>
>> **"task"** string <br>
>> **"due"** string <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Input:

```json
{
    "id":1,
    "task":"Pay pomos",
    "due":"Friday",
    "complete":true
}
```
 
##### Output Fields:

>> **"ID"** number :warning:Not currently used <br>
>> **"CreatedAt"** string :warning:Not currently used <br>
>> **"UpdatedAt"** string :warning:Not currently used <br>
>> **"DeletedAt"** string :warning:Not currently used <br>
>> **"id"** number <br>
>> **"task"** string <br>
>> **"due"** string <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
{
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "id": 1,
    "task": "Pay pomos",
    "due": "Friday",
    "complete": true
}
```

#### Delete

> DELETE 1270.0.1:8081/todo/{id} <br>

Deletes the Todo items with the id designated by the address, returns the deleted Todo item<br>

##### Output Fields:

>> **"ID"** number :warning:Not currently used <br>
>> **"CreatedAt"** string :warning:Not currently used <br>
>> **"UpdatedAt"** string :warning:Not currently used <br>
>> **"DeletedAt"** string :warning:Not currently used <br>
>> **"id"** number <br>
>> **"task"** string <br>
>> **"due"** string <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
{
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "id": 0,
    "task": "",
    "due": "",
    "complete": false
}
```

### Login

*Note yet completed*
