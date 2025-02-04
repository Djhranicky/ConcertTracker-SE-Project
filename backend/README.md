# Backend

To start a local development server, run:

```bash
make run
```

Once the server is running, test requests can be made to `localhost:8080` using postman or other API testing software

## Functionality

The API is started up through cmd/main.go. This sets up the database connection and the http server.

When the http server is ran, a new router is created which sets up different routes and functions to handle requests to those routes.
