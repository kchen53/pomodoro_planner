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
    - [Signup](#signup) <br>
    - [Login](#login) <br>
    - [Logout](#logout) <br>
    - [Get Username](#get-username) <br>

---

### Todo

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
        "repeat": 32, //Note: Every Friday
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
    "repeat": 1, //Note: Every Sunday
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

### Login

####  Signup

> <font color="orange">POST</font> 127.0.0.1:8081/user <br>

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

> <font color="blue">PUT</font> 127.0.0.1:8081/user <br>

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

> <font color="red">DELETE</font> 127.0.0.1:8081/user <br>

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

> <font color="green">GET</font> 127.0.0.1:8081/user <br>

Returns the username of the currently logged in user, returns "Anonymous" if not logged in; returns "" if there is a server side error <br>
 
##### Output Fields:

>> string <br>
> 
> Style: Raw JSON

###### Example Output:

```json
"Anonymous"
```
          

