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

`"/register"`

`"/login"`

`"/validate"`

`"/swagger"`
