# ConcertTracker-SE-Project

## Description

There are existing technologies to share your interest in hobbies (letterboxd, goodreads, etc.), but none exist for concerts. User will be able to login, mark a concert as attended, add reviews, and follow other users connected with their spotify account. We plan to integrate with the Spotify API to keep track of your friend network and see upcoming concerts in your area. We will use the setlist.fm API to pull a historical record of concerts so users can mark past events they have been to. Future features will be added as necessary.


## Contributers

### Front End
- Srija Kethireddy
- Paola Solari

### Back End
- David Hranicky
- Parth Shah

## Tech Stack

### Front End
- Angular 
### Back End
- Golang

# Backend

To start a local development server, run:

```bash
make run
```

Once the server is running, test requests can be made to `localhost:8080` using postman or other API testing software

To update swagger documentation, run:
1. Add comments like this before corresponding function in `routes.go`:
```ts
// @Summary Register user
// @Description Registers a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body types.UserRegisterPayload true "Register Payload"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {string} string "Invalid payload or user already exists"
// @Router /register [post]
```
2. Run:
```bash
swag init -g cmd/main.go
```

Note: for linux, if `swag` command is not recognised, make sure to set your path using:
```bash
export PATH=$(go env GOPATH)/bin:$PATH
```

## Functionality

The API is started up through cmd/main.go. This sets up the database connection and the http server.

When the http server is run, a new router is created which sets up different routes and functions to handle requests to those routes.

## References/Documentation

Useful references can be listed here for future reference:

* Potential session management options:
  * [Cookies](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.2.html)
  * [Middlewares](https://medium.com/@fasgolangdev/how-to-create-a-secure-authentication-api-in-golang-using-middlewares-6988632ddfd3)
  * [Pre-made package](https://github.com/alexedwards/scs)
  * [API Keys](https://dev.to/caiorcferreira/implementing-a-safe-and-sound-api-key-authorization-middleware-in-go-3g2c)

# Frontend

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 19.1.4.

## Development server

To start a local development server, run:

```bash
ng serve
```

Once the server is running, open your browser and navigate to `http://localhost:4200/`. The application will automatically reload whenever you modify any of the source files.

## Code scaffolding

Angular CLI includes powerful code scaffolding tools. To generate a new component, run:

```bash
ng generate component component-name
```

For a complete list of available schematics (such as `components`, `directives`, or `pipes`), run:

```bash
ng generate --help
```

## Building

To build the project run:

```bash
ng build
```

This will compile your project and store the build artifacts in the `dist/` directory. By default, the production build optimizes your application for performance and speed.

## Running unit tests

To execute unit tests with the [Karma](https://karma-runner.github.io) test runner, use the following command:

```bash
ng test
```

## Running end-to-end tests

For end-to-end (e2e) testing, run:

```bash
ng e2e
```

Angular CLI does not come with an end-to-end testing framework by default. You can choose one that suits your needs.

## Additional Resources

For more information on using the Angular CLI, including detailed command references, visit the [Angular CLI Overview and Command Reference](https://angular.dev/tools/cli) page.
