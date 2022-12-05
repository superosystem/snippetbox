# Snippetbox

## Routing request
| Method | URL Pattern | Handler    | Action    |
| :---:   | :---: | :---: | :---: |
| GET | /   | home   | Display the home page   |
| GET | /snippet/view?id=1    | snipperView   | Display a specific snippet   |
| GET | /snippet/create   | snippetCreate   | Display a HTML form for creating a new snipper   |
| POST | /snippet/create   | snippetCreatePost   | Create a new snippet   |
| GET | /users/signup   | userSignup   | Display a HTML form for signing up a new user   |
| POST | /users/signup   | userSignupPost   | Create a new user   |
| GET | /users/login   | userLogin   | Display a HTML form for logging in a user   |
| POST | /users/login   | userLoginPost   | Authenticate and login the user   |
| POST | /users/logout   | userLogoutPost   | Logout the user   |
| GET | /static/*filepath   | http.FileServer   | Server a specific static file   |
| GET | /ping   | ping   | Return a 200 OK response   |


## Security Improvements

* Make LTS with running generate_cert.go
```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

## Courses
- Foundations
    - Setup a project directory which follows the Go conventions.
    - Start a web server and listen for incoming HTTP requests.
    - Route requests to different handlers based on the request path.
    - Send different HTTP responses, headers and status codes to users.
    - Fetch and validate untrusted user input from URL query string parameters.
    - Structure your project in a sensible and scalable way.
    - Render HTML pages and use template inheritance to keep your markup free of duplicate boilerplate code.
    - Serve static files like images, CSS and JavaScript from your application.
-  Configuration and Error Handling
    - Set configuration settings for your application at runtime in an easy and idiomatic way using command-line flags.
    - Improve your application log messages to include more information, and manage them differently depending on the type (or level) of log message.
    - Make dependencies available to your handlers in a way that’s extensible, type-safe, and doesn’t get in the way when it comes to writing tests.
    - Centralize error handling so that you don’t need to repeat yourself when writing code.
-  Database Driven Responses
    - Install a database driver to act as a ‘middleman’ between MySQL and your Go application. 
    - Connect to MySQL from your web application (specifically, you’ll learn how to establish a pool of reusable connections).
    - Create a standalone models package, so that your database logic is reusable and decoupled from your web application.
    - Use the appropriate functions in Go’s database/sql package to execute different types of SQL statements, and how to avoid common errors that can lead to your server running out of resources.
    - Prevent SQL injection attacks by correctly using placeholder parameters.
    - Use transactions, so that you can execute multiple SQL statements in one atomic action.
-  Dynamic HTML Templates
    - Pass dynamic data to your HTML templates in a simple, scalable and type-safe way.
    - Use the various actions and functions in Go’s html/template package to control the display of dynamic data.
    - Create a template cache so that your templates aren’t being read from disk for each HTTP request. 
    - Gracefully handle template rendering errors at runtime. 
    - Implement a pattern for passing common dynamic data to your web pages without repeating code. 
    - Create your own custom functions to format and display data in your HTML templates.
-  Middleware
    - An idiomatic pattern for building and using custom middleware which is compatible with net/http and many third-party packages.
    - How to create middleware which sets useful security headers on every HTTP response.
    - How to create middleware which logs the requests received by your application.
    - How to create middleware which recovers panics so that they are gracefully handled by your application.
    - How to create and use composable middleware chains to help manage and organize your middleware.
-  Advanced Routing
    - Briefly discuss the features of a few good third-party routers.
    - Update our application to use one of these routers and demonstrate how to use method-based routing and clean URLs.
-  Processing Forms
    - How to parse and access form data sent in a POST request.
    - Some techniques for performing common validation checks on the form data.
    - A user-friendly pattern for alerting the user to validation failures and re-populating formfields with previously submitted data.
    - How to keep your handlers clean by using helpers for form processing and validation.
-  Stateful HTTP
    - What session managers are available to help us implement sessions in Go.
    - How to use sessions to safely and securely share data between requests for a particular user.
    - How you can customize session behavior (including timeouts and cookie settings) based on your application’s needs.
-  Securiy Improvements
    - How to quickly and easily create a self-signed TLS certificate, using only Go.
    - The fundamentals of setting up your application so that all requests and responses are served securely over HTTPS.
    - Some sensible tweaks to the default TLS settings to help keep user information secure and our server performing quickly.
    - How to set connection timeouts on our server to mitigate slow-client attacks.
-  User Authentication
    - How to implement basic signup, login and logout functionality for users.
    - A secure approach to encrypting and storing user passwords securely in your database using Bcrypt.
    - A solid and straightforward approach to verifying that a user is logged in using middleware and sessions.
    - How to prevent cross-site request forgery (CSRF) attacks.
-  Using Request Context
    - What request context is, how to use it, and when it is appropriate to use it.
    - How to use request context in practice to pass information about the current user between your handlers.
-  Optional GO Feature
    - File embedding makes it possible to embed external files into your Go program itself.
    - Generics can help to reduce the amount of boilerplate code you need to write, while retaining compile-time type safety.
-  Testing
    - How to create and run table-driven unit tests and sub-tests in Go. 
    - How to unit test your HTTP handlers and middleware. 
    - How to perform ‘end-to-end’ testing of your web application routes, middleware and handlers.
    - How to create mocks of your database models and use them in unit tests. 
    - A pattern for testing CSRF-protected HTML form submissions.
    - How to use a test instance of MySQL to perform integration tests.
    - How to easily calculate and profile code coverage for your tests.