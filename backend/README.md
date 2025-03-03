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
