# Sprint 2

## Frontend

### Work completed

- [x] [Created the frontend for the User-Profile page.](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/27) This page can only be accessed if the user is logged in. The page is currently populated with mock user data.
- [x] [Created High Fidelity wireframes for Concert, Dashboard, and all User Page's views](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/3)
- [x] [Connected backend to login/registration pages](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/30) Created an auth service that connects the frontend login/registration as well as session management to the backend.
- [x] [Created a Not Found Page to be served whenever uses goes to a Route it doesn't exist.](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/26)
- [x] [Implemented form validation for both login and register pages.](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/43). This includes checking if user submitted all required fields, if email is valid, if password match, and if password is at least 6 characters long. If any of these conditions are not met, form should not be submitted.
- [x] Added logout functionality from the frontend. Clears localStorage of isAuth value from frontend.
- [x] [Modified Navbar to change contents once user is logged in](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/20). Logged In user sees User avatar and log out button, unsigned user sees log in and sign up buttons.
- [x] Implemented Route Guards to only show certain pages if user is logged in or not.
- [x] [Started Dashboard Page UI](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/53)
- [x] [Created Posts injectable service that currently serves mock data of posts for Dashboard page, allowing for a seamless connection to backend data.](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/51)
- [x] [Created Post modular component to be used in Dashboard, Concert and Uses Pages](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/46)
- [x] [Created Cypress E2E tests for login and register pages](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/46)
- [x] [Created Jasmine unit tests for existing components including landing page subcomponents, existing services and login and register pages.](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/44)


### Frontend Unit Tests
#### Jasmine Unit Tests
- App Component
  - should be created 
- NavBar
  - should be created   
- Landing Page
  - should be created
  - *Landing Page Components*
  - Discover Component
    - should be created
  - Discuss Component
    - should be created
    - should bind title
    - should bind subtitle
    - should bind paragraph
  - Features Component
    - should be created
  - Hero
    - should be created
    - should have a button with routerLink set to /register
    - should have a button with routerLink set to /login
    - should navigate to register page when clicking sign up button
    - should navigate to register page when clicking log in button
  - Popular Concerts Carousel
    - should be created
  - Tweets / What Others are Saying
    - should be created
- Login Page
  - should be created
  - should render username and password input fields and login button
  - should have email and passwords as required fields of form
  - should require email field
  - should validate email field to be email format
  - should require password field
  - should call login() if form valid
  - should not call AuthenticationService.login if form is invalid
- Registration Page
  - should be created
  - should render username and password input fields and login button
  - should have email, user and passwords as required fields of form
  - should require email field
  - should validate email field to be email format
  - should require password field
  - should not call AuthenticationService.register if form is invalid
- Not Found Page
  - should be created
- User Profile Page
  - should be created
  - should display the correct user name
  - should have "profile" as the default active tab
  - should change active tab when a tab is clicked
  - should display the correct number of favorite concerts
  - should display the correct statistics text
  - should display the correct number of recent activities
  - should display the bucket list section
- Authentication Service
  - should be created
  - should send a POST to /register endpoint in backend
  - should handle existing user register error
  - should send a POST to /login endpoint in backend
  - should handle existing user login error
  - should delete session from localStorage
  - should return true if session exists in localStorage
  - should return false if session does not exist in localStorage
- Pop Concerts Service
  - should be created
  - should return an observable of tours
  - should return an observable of the correct type

#### Cypress E2E Tests
- Login Page
  - Visits the login page
  - should render username and password input fields and login button
  - should have email and passwords as required fields of form
  - should show validation error for invalid email format
  - should submit form and navigate to dashboard on successful login
  - should redirect to /register when clicking register on description
- Registration Page
  - Visits the Register page
  - should render email, username and password input fields and sign up button
  - should have email, user and passwords as required fields of form
  - should show validation error for invalid email format
  - should submit form and navigate to login page on successful registration
  - should redirect to /login when clicking log in here on description
- Spec Test
  - Visits inital project Page 
## Backend

### Work completed

Added session management by storing a JWT string in the browser cookie. This can be checked in the future to ensure that a user's "session" is still active and does not need to log back in.

Added swagger documentation for existing routes. Allows front end team to build against expected functionality and ensures information is shared consistently.

Began investigating utilizing the setlist.fm api for concert information. Applied and received API key to be used and started building understanding of what data will be needed from the API.

Began architecting needed database tables. Deciding what information needs to be stored and how best to design database to accomodate application needs.

### Unit Tests

Wrote unit tests for each route implemented so far and for managing the session with JWT tokens stored in cookies

Unit tests written so far:

- TestUserServiceHandleRegister
  - Should fail if request body is empty
  - Should fail if the user payload is invalid
  - Should fail if user exists
  - Should succeed when new user is created
- TestUserServiceHandleLogin
  - Should fail if request body is empty
  - Should fail if payload is invalid
  - Should fail if user does not exist
  - Should fail if user enters wrong password
  - Should pass if user enters correct user name and password
- TestUserServiceHandleValidate
  - should fail when no id cookie is present
  - should fail when invalid jwt string is present
  - should pass when valid cookie is present
- TestSessionMethods
  - should fail if request has no cookie
  - should pass if request has cookie
  - verification should fail if no cookie present
  - verification should fail if no JWT token present
  - verification should fail if JWT token is expired
  - verification should succeed if JWT token is valid

### API Documentation

The API currently consists of five endpoints mostly focused on authentication functionality.

#### Home Route

The `"/"` endpoint (home route) is a basic endpoint that currently only returns a status code 200 and a hello world message. This may be adapted for later use to serve information to the landing page if necessary, and mostly exists to verify the API can serve responses.

#### Register Route

The `"/register"` endpoint is used to register a new user in the system. Users currently have a name, email and password stored in the database. The JSON payload expected looks like

```json
{
  "Name":"Name",
  "Email":"test@example.com",
  "Password":"Password",
}
```

The payload is first validated to ensure that the inputs meet the expected format. Then if the user does not already exist in the system, a new user is created in the database and the endpoint returns a 201 code indicating the account was created successfully.

#### Login Route

The `"/login"` endpoint is used to authenticate a user and create a session. The JSON payload expected looks like

```json
{
  "Email":"test@example.com",
  "Password":"Password",
}
```

The endpiont first verifies that a valid email and password combination was supplied. This returns a 400 code with an `"invalid email or password"` message if the credentials are not valid. This then creates a new JSON Web Token that encodes a userId number and an expiration time with a private secret key. Finally, this JWT is stored as a cookie in the response and a status code of 200 is returned.

#### Validate Route

The `"/validate"`endpoint is used to check if a user currently has a valid session. No payload is consumed by this endpoint. This attempts to retrieve the expected cookie provided by the request. If no cookie with the expected name is present or the contained token is invalid, then a 401 code is returned. If the cookie contains a valid token, then a 200 code is returned.

#### Swagger Route

The `"/swagger"` endpoint is used to document and test the functionality of the backend API. This route is automatically generated by the `github.com/swaggo/swag` package and allows for test requests to be sent to the backend API.
