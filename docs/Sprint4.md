# Sprint 4

## Frontend

### Work completed

- [x] [Created the frontend for the Concert page.](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/28) The page is currently populated with mock concert data.
- [x] [Refined Post UI modular component to be used in Dashboard, Concert and User Pages](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/53). Currently serves mock post data and the toggle like functionality is static.
- [x] [Created Dashboard Page UI](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/54)
Currently serves mock posts from posts service. Auth guarded so only logged in users should reach it when navigating to '/'
- [x] [Created injectable service for User Page](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/36) To be modified to connect with backend once endpoints are set.
- [x] [Created mock data for User Page ](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/32) Created interfaces for user service.
 - [x] [Connected frontend to the /validate endpoint in backend](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/55) Modified isAuth() function in authenticationService to check if the jwt given when logged in is still valid. 
 - [x] [Created injectable service for Concert and Artist Page](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/33) To be modified to connect with backend once endpoints are set. Current functions getConcert(), getArtist(), getRecentConcerts() and getUpcomingConcerts()
- [x] [Improved service for serving Posts](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/51)
- [x] [Created mock data for Concert and Artist Page](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/34) Created and refactored interfaces for concert and artist services.
- [x] [Added navigation menu to signed in Navbar](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/58) Created and refactored interfaces for concert and artist services. Moved sign out button to drop down menu bar on user avatar click.
- [x] [Refactored User Page to use modular post components.](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/58) Added 'Activity', 'Concerts', and 'Lists' Tabs Views for existing user Page.
- [x] [Created Cypress E2E tests for Artist, Concert and User pages](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/46)
- [x] [Created Jasmine unit tests for new pages and components](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/44)
- [x] [Initial Search functionality implemented and connected to backend](https://github.com/Djhranicky/ConcertTracker-SE-Project/issues/66) Created SearchResults page that can be accessed through the search bar, which displays results from the search /artist endpoint from the backend.
- [x] Created Home Component to render landing page if user not logged in and the dashboard if logged in when navigating to '/' 

### Frontend Unit Tests
#### Jasmine Unit Tests
- App Component
  - should be created
- HomeComponent
  - should create
  - should set isLoggedIn false when auth invalid
  - should set isLoggedIn true when auth valid
- NavBar
  - should show menubar items if logged in
  - should log out if logged in
  - should have a button with routerLink set to /login when logged out
  - should not show menubar items if not logged in
  - should have a button with routerLink set to /register when logged out
  - should navigate to register page when clicking log in button when logged out
  - should navigate to register page when clicking sign up button when logged out
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
  - should call AuthenticationService.login if form is valid
  - should not call AuthenticationService.login if form is invalid
- Registration Page
  - should be created
  - should render username and password input fields and login button
  - should have email, user and passwords as required fields of form
  - should require email field
  - should validate email field to be email format
  - should require password field
  - should not call AuthenticationService.register if form is invalid
  - should call AuthenticationService.register if form is valid
  - should handle registration error
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
  - should return true if /validate return 200 message
  - should return false if /validate returns 401
  - should handle existing user register error
  - should send a POST to /login endpoint in backend
  - should handle existing user login error
  - should delete cookie
  - should be created
  - should send a POST to /register endpoint in backend
- Pop Concerts Service
  - should be created
  - should return an observable of tours
  - should return an observable of the correct type
- ConcertComponent
  - should create
- SearchComponent
  - should create
- ArtistComponent
  - should create
- PostService
  - should return an observable of posts
  - should return an observable of the correct type
  - should be created
- PostComponent
  - should create
  - should toggle isLiked affect post like count
- ConcertService
  - should be created
- DashboardComponent
  - should create
- UserService
  - should convert concert card to post format
  - should return user posts
  - should return user profile
  - should return favorite concerts
  - should be created
- TimeAgoPipe
create an instance
- SearchPage
   - should create
   - should contain navbar and search components
- SearchComponent
  - should create
  - should show error message when error exists and not loading
  - should get query parameter and search for artist on init
  - should initialize results array when there is no query
  - should display search results when data is loaded
  - should set error message when HTTP request fails
  - should make HTTP request with correct URL when searching


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
- Artist Page
  - Visits artists Page
  - should display artist details
  - should display recent shows
  - should display upcoming shows
  - should display artist stats
- Concert Page
  - Visits Concerts Page 
  - should display concert header
  - should mark concert as attended
  - should display the setlist
  - should display recent activity
  - should display Spotify playlist
  - should display attended users
  - should display other shows and festivals
- User Page
  - Visits user-profile page
  - should display user profile information correctly
  - should navigate between tabs correctly
  - should display favorite concerts on profile tab
  - should display posts on the activity tab
  - should display posts on the concerts tab
  - should display edit profile icon and handle image upload
  - should handle responsive layout
- Spec Test
  - Visits inital project Page 
## Backend

### Work completed

Add to database structure to capture many elements required for user interaction with the site.

Create /userpost endpoint to allow users to create posts to the site. The POST route will create a new post. The GET route will return a list of posts to populate a user's followed network.

Create /like endpoint to allow users to like/unlike a post. The POST route will toggle whether a post is liked by a given user. The GET route will return the number of likes for a given post.

Create /follow endpoint to allow users to follow other users. The POST route will toggle whether a user follows another user. The GET route takes two options as a query parameter. type = followers will give a list of users that follow the given user. type = following will give a list of users that a given user follows.

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
- TestArtistServiceHandleArtist
  - should pass if artist already in database
  - should pass if artist found in external API
- TestArtistServiceHandleImport
  - should fail with no name query parameter
  - should fail with invalid full query parameter
  - should fail if artist mbid not in database
  - should fail if artist mbid not in external API
  - should pass if artist mbid in database
  - should pass if artist mbid in database for full import
- TestConcertServiceHandleConcert
  - should fail with no id query parameter
  - should fail if setlist not found in external API
- TestUserServiceHandlePost
  - should fail with no 'authorID' field included in incoming json
  - should fail with no 'type' field included in incoming json
  - should fail with no 'isPublic' field included in incoming json
  - should fail with no 'concertID' field included in incoming json
  - should fail if invalid 'type' supplied
  - should pass and create post in database
  - GET should fail with no userID query parameter
  - GET should fail with bad userID query parameter
  - GET should pass with valid userID query parameter
- TestUserServiceHandleLike
  - should fail if UserPostID not included
  - should fail if UserID not included
  - should succeed when user first likes a post
  - should succeed when user removes like from post
  - should succeed when user likes a post again
  - GET should fail when query param not provided
  - GET should fail when bad query param provided
  - GET should return like count for valid input
- TestUserServiceHandleFollow
  - should fail if FollowedUserID not included
  - should fail if UserID not included
  - should succeed when user first follows another user
  - should succeed when user unfollows another user
  - should succeed when user follows another user again
  - GET should fail if userID param not included
  - GET should fail if type param not included
  - GET should fail if type param invalid
  - GET should fail bad userID provided
  - GET should pass with valid userID and type = followers
  - GET should pass with valid userID and type = following
- TestUserInfoRoute
  - should fail with no username in payload
  - should fail with empty username
  - should fail with non-existent username
  - should succeed with valid username
  - should handle OPTIONS request correctly
- TestUsersRoute
  - should return current users list
  - should return updated list after adding new users
  - should handle OPTIONS request correctly

### API Documentation

The API currently consists of eight endpoints.

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

#### Artist Route

The `"/artist"` endpoint retrieves information about a specific artist by searching setlist.fm. It accepts an artist's name as input and returns details such as the artist's name, recent setlists, and past concert performances. Additionally, it provides the total number of setlists available, associated tour names, and any upcoming shows (if applicable) listed on setlist.fm. The setlists are returned in reverse chronological order.

#### Import Route

The `"/import"` endpoint is used to fetch and consume information from the external setlist.fm API. No payload is consumed by this endpoint. This endpoint parses concert information returned from the API call and creates data in the database, including records for additional artists, venues, tours, concerts, songs, and the concert-song relation. A query parameter of the artist MBID is required, and an optional full parameter can be specified. With no full parameter, this will only import up to 20 concerts for an artist. With the full parameter set to true, this will import all concerts for a given artist. This returns a 201 code on success, and a 400 or 504 code on failure.

#### Concert Route

The `"/concert"` endpoint fetches details about a specific concert using its unique concert ID from the setlist.fm API. It returns comprehensive information, including the setlist of songs performed, the associated tour, and venue details such as the venue name, city, and country.

#### UserPost Route

The `"/userpost"` endpoint created posts into the database.

For the POST route, the JSON body must include the AuthorID, the Type, the ConcertID, and whether or not the post IsPublic. The body can optionally include a review, a rating, and a UserPostID if the post being created is a comment to another post. The type can only be one of ATTENDED, LISTCREATED, WISHLIST, or REVIEW. This data will be created as a post in the UserPost table.

For the GET route, a userID query parameter must be passed in, and an optional page number parameter can be included. This returns a list of public posts created by the user with the given userID.

#### Like Route

The `"/like"` endpoint is used to create and fetch likes on posts.

The POST route takes a UserID and a UserPostID to add a like from the user to the post. Calling the post route with a record that already exists will remove the like from the table.

The GET route returns the number of likes for a given UserPostID.

#### Follow Route

The `"/follow"` endpoint allows users to follow eachother and can return lists of users.

The POST route will add or remove a follow relation between two users. When the UserID and FollowedUserID pairing is not present, a record is created. If the pairing is present, the record is removed (unfollowing).

The GET route will return a list of users in two different scenarios. The type = followers scenario will return a list of users that all follow the given user, and the type = following will return a list of users that the given user follows.
