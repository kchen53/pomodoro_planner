# Sprint 3

## Work Completed
### Frontend:
* Timer funtionality implemented
* Login mostly completed
* Calendar now displays
* Deafult home page
* Day events display on the session page
* Added spotify playist on session page 

### Backend:

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

it('check the to-do list html header', () => {
  fixture.detectChanges();
  const compiled = fixture.nativeElement;
  expect(compiled.querySelector('mat-card-title').textContent).toContain('To-Do List');
});

##### New Tests:
it('check if the calendar displays the correct month', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('h2').textContent).toContain('March 2023');
  });
  
it('check if the daily calender displays the correct date', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('h2').textContent).toContain('Wednesday, March 29, 2023');
  });
  
it('should have a link to the session page', () => {
  const link = fixture.debugElement.query(By.css('a[routerLink="session"]'));
  expect(link).toBeTruthy();
});

it('should have a settings icon', () => {
  const icon = fixture.debugElement.query(By.css('mat-icon'));
  expect(icon.nativeElement.innerText).toContain('settings');
});

it('check if the component loads', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('mat-card-content').textContent).toContain('Spotify Api');
  });
  
it('check the to-do list html header', () => {
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('h2').textContent).toContain('To-Do List');
  });

it('should have an app-calendar component', () => {
    const appCalendar = fixture.debugElement.query(
      By.directive(CalendarComponent)
    );
    expect(appCalendar).not.toBeNull();
  });
 
 it('should have an app-dailyTask component', () => {
    const appCalendar = fixture.debugElement.query(
      By.directive(DailyTaskComponent)
    );
    expect(appCalendar).not.toBeNull();
  });
  
#### Cypress:

describe('Login', () => {
  beforeEach(() => {
    cy.visit('http://localhost:8081/home')
  })
  it('passes', () => {
    cy.contains('Login').click()
    cy.get('input[name="username"]').type('Username')
    cy.get('input[name="password"]').type('password')
    cy.get('button').contains('Submit').click()

    cy.contains('Login').click()
    cy.contains('Sign Up').click()
    cy.get('input[name="username"]').type('Username')
    cy.get('input[name="password"]').type('password')
    cy.get('button').contains('Submit').click()
  })
})

describe('Timer', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:8081/session')
    cy.get('input[placeholder="mm:ss"]').type('00:20')
    cy.get('button').contains('Start').click()
    cy.get('div[class="timer-value"]').contains('00:20')
    cy.wait(20*1000)
    cy.get('div[class="timer-value"]').contains('00:00')
  })
})

describe('Todo', () => {
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

describe('Calendar', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:8081/')

    cy.get('app-calendar').should('be.visible')

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
- [**User**](#user)
    - [Signup](#signup) <br>
    - [Login](#login) <br>
    - [Logout](#logout) <br>
    - [Get Username](#get-username) <br>

---

### Todo

##### Fields:

> **"id"** number <br>
> Unique key for every todo. Used for updating, deleting, or getting specific todo items

> **"name"** string <br>
> String name for todo describing the task

> **"date"** string <br>
> Due date is held in a string of format "YYYY-MM-DD" where YYYY is the year, MM is the month, DD is the day

> **"time"** number *Format: seconds (null OK)* <br>
> Integer representing time spent/time left for task. Time is in seconds and can be null if not needed.

> **"repeat"** number <br>
> Signifies how often to repeat task. Data is the integer value of a 7 digit binary flags: <br>
>> | Sun | Mon | Tue | Wed | Thu | Fri | Sat |
>> |-----|-----|-----|-----|-----|-----|-----|
>> |  0  |  0  |  0  |  0  |  0  |  0  |  0  |
> Example 1. If todo is to repeat every Sunday, then the value would be bin(1000000) = int(64) <br>
> Example 2. If todo is to repeat every Monday, Wednesday, and Friday, then the value would be bin(0101010) = int(42) <br>

> **"complete"** boolean <br>
> Boolean value signifying whether the task has been completed or not

####  Create

> <font color="orange">POST</font> 127.0.0.1:8081/todo <br>

Creates a new Todo item, returns the created Todo item<br>

##### Input Fields:

>> **"name"** number <br>
>> **"date"** string *Format: "YYYY-MM-DD"* <br>
>> **"time"** number *Format: seconds (null OK)* <br>
>> **"repeat"** number *Format: Binary flags (see Todo)* <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Input:

```json
{
    "name":"Pay pomos",
    "date":"2021-03-26",
    "time":3600,
    "repeat":0,
    "complete":false
}
```
 
##### Output Fields:

>> **"id"** number <br>
>> **"name"** string <br>
>> **"date"** string *Format: "YYYY-MM-DD"* <br>
>> **"time"** number *Format: seconds (null OK)* <br>
>> **"repeat"** number *Format: Binary flags (see Todo)* <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
{
    "id": 0,
    "name": "Pay pomos",
    "date": "2021-03-26",
    "time": 3600,
    "repeat": 0,
    "complete": false
}
```
          
#### Get All

> GET 127.0.0.1:8081/todo <br>

Returns a list of all stored Todo items <br>

##### Output Fields:

>> **"id"** number <br>
>> **"name"** string <br>
>> **"date"** string *Format: "YYYY-MM-DD"* <br>
>> **"time"** number *Format: seconds (null OK)* <br>
>> **"repeat"** number *Format: Binary flags (see Todo)* <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
[
    {
        "id": 0,
        "name": "Pay pomos",
        "date": "2021-03-26",
        "time": 3600,
        "repeat": 0,
        "complete": true
    },
    {
        "id": 2,
        "name": "Release doros",
        "date": "2021-03-31",
        "time": null,
        "repeat": 2, //Note: Every Friday
        "complete": false
    }
]
```

#### Get by ID

> GET 127.0.0.1:8081/todo/{id} <br>

Returns the Todo item with the id designated by the address

##### Output Fields:

>> **"id"** number <br>
>> **"name"** string <br>
>> **"date"** string *Format: "YYYY-MM-DD"* <br>
>> **"time"** number *Format: seconds (null OK)* <br>
>> **"repeat"** number *Format: Binary flags (see Todo)* <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
{
    "id": 0,
    "name": "Pay pomos",
    "date": "2021-03-26",
    "time": 3600,
    "repeat": 64, //Note: Every Sunday
    "complete": true
}
```
#### Update

> PUT 127.0.0.1:8081/todo/{id} <br>

Modifies the Todo item with the id designated by the address, returns the updated Todo item<br>

##### Input Fields:

##### Input Fields:

>> **"name"** number <br>
>> **"date"** string *Format: "YYYY-MM-DD"* <br>
>> **"time"** number *Format: seconds (null OK)* <br>
>> **"repeat"** number *Format: Binary flags (see Todo)* <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Input:

```json
{
    "name":"Pay pomos",
    "date":"2021-03-26",
    "time":3600,
    "repeat":0,
    "complete":true
}
```
 
##### Output Fields:

>> **"id"** number <br>
>> **"name"** string <br>
>> **"date"** string *Format: "YYYY-MM-DD"* <br>
>> **"time"** number *Format: seconds (null OK)* <br>
>> **"repeat"** number *Format: Binary flags (see Todo)* <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
{
    "id": 0,
    "name": "Pay pomos",
    "date": "2021-03-26",
    "time": 3600,
    "repeat": 0,
    "complete": true
}
```

#### Delete

> DELETE 1270.0.1:8081/todo/{id} <br>

Deletes the Todo items with the id designated by the address, returns the deleted Todo item<br>

##### Output Fields:

>> **"id"** number <br>
>> **"name"** string <br>
>> **"date"** string *Format: "YYYY-MM-DD"* <br>
>> **"time"** number *Format: seconds (null OK)* <br>
>> **"repeat"** number *Format: Binary flags (see Todo)* <br>
>> **"complete"** boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
{
    "id": 7,
    "name": "",
    "date": "",
    "time": 0,
    "repeat": 0,
    "complete": false
}
```

### User

####  Signup

> <font color="orange">POST</font> 127.0.0.1:8081/signup <br>

If the username does not already exist then creates a new user and logs in, returns whether a new user was successfully created <br>

##### Input Fields:

>> **"username"** string <br>
>> **"password"** string <br>
> 
> Style: Raw JSON

###### Example Input:

```json
{
    "username":"Pomodoro",
    "password":"ILovePomos",
}
```
 
##### Output Fields:

>> boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
true
```

####  Login

> <font color="blue">PUT</font> 127.0.0.1:8081/login <br>

Logs in if a the a user with the given credentials exists, returns whether the login was successful <br>

##### Input Fields:

>> **"username"** string <br>
>> **"password"** string <br>
> 
> Style: Raw JSON

###### Example Input:

```json
{
    "username":"Pomodoro",
    "password":"ILovePomos",
}
```
 
##### Output Fields:

>> boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
false
```

####  Logout

> <font color="red">PUT</font> 127.0.0.1:8081/logout <br>

Logs out, setting the user to "Anonymous"<br>
 
##### Output Fields:

>> boolean <br>
> 
> Style: Raw JSON

###### Example Output:

```json
true
```
####  GetUsername

> <font color="green">GET</font> 127.0.0.1:8081/userdata <br>

Returns the username of the currently logged in user, returns "Anonymous" if not logged in; returns "" if there is a server side error <br>
 
##### Output Fields:

>> string <br>
> 
> Style: Raw JSON

###### Example Output:

```json
"Anonymous"
```
