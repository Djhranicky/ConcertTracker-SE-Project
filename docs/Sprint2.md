# Sprint 2

## Frontend

### Frontend Unit Tests

## Backend

### Work completed

Added session management by storing a JWT string in the browser cookie. This can be checked in the future to ensure that a user's "session" is still active and does not need to log back in.

Added swagger documentation for existing routes. Allows front end team to build against expected functionality and ensures information is shared consistently.

Began investigating utilizing the setlist.fm api for concert information. Applied and received API key to be used and started building understanding of what data will be needed from the API.

Began architecting needed database tables. Deciding what information needs to be stored and how best to design database to accomodate application needs.

### Unit Tests

Wrote unit tests for each route implemented so far and for managing the session with JWT tokens stored in cookies

Unit tests written so far:

- Testing handleRegister
  - Should fail if request body is empty
  - Should fail if the user payload is invalid
  - Should fail if user exists
  - Should succeed when new user is created
- Testing handleLogin
  - Should fail if request body is empty
  - Should fail if payload is invalid
  - Should fail if user does not exist
  - Should fail if user enters wrong password
  - Should pass if user enters correct user name and password

### API Documentation

The API currently consists of five endpoints mostly focused on authentication functionality.

The `"/"` endpoint (home route) is a basic endpoint that currently only returns a status code 200 and a hello world message. This may be adapted for later use to serve information to the landing page if necessary, and mostly exists to verify the API can serve responses.

The `"/register"` endpoint is used to register a new user in the system. Users currently have a name, email and password stored in the database. The JSON payload expected looks like

```json
{
  "Name":"Name",
  "Email":"test@example.com",
  "Password":"Password",
}
```

The payload is first validated to ensure that the inputs meet the expected format. Then if the user does not already exist in the system, a new user is created in the database and the endpoint returns a 201 code indicating the account was created successfully.

The `"/login"` endpoint is used to authenticate a user and create a session. The JSON payload expected looks like

```json
{
  "Email":"test@example.com",
  "Password":"Password",
}
```

The endpiont first verifies that a valid email and password combination was supplied. This returns a 400 code with an `"invalid email or password"` message if the credentials are not valid. This then creates a new JSON Web Token that encodes a userId number and an expiration time with a private secret key. Finally, this JWT is stored as a cookie in the response and a status code of 200 is returned.

The `"/validate"`endpoint

`"/swagger"`
