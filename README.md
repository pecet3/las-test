# PDF - LAS

For this project I used the latest version of Go and React Vite.
I focused more on backend than on frontend so frontend can be messy and ugly in some parts.

Here You can find: [diagram](https://drive.google.com/file/d/1y6VKN10qqk3eMB7XUsYxdj1WsIE58viI/view?usp=sharing)

Let me know if You have any questions :)

## Backend - Design

I like standard library in Go, so I didn't use any router such as Gin etc.
Backend is divided into two major parts - cmd and pkg. In cmd dir, there is implementation of packages, so routing is there and business logic. In pkg dir there are packages used in cmd.

Repository pattern is used for dependency injection to keep code clean and easy to maintenance.

Used packages and technologies:

- Sqlite3 and sqlc as good replacement of ORM
- UUID package from google
- Godotenv to read easy .env file
- JWT package for auth
- Validator V10, beacuse We cannot never trust user's input
- My own logger package
- My own passwordless auth package, but it is directly in pkg directory, not as dependency

## Frontend - Design

Frontend is not fancy as I said, but there are essentials features such as:

- View on the site (modal which is using iframe html tag)
- Context and state managment
- Protected Page Wrapper which check if user is logged: sending request to /api/auth/ping. This endpoint checks cookies if contain valid JWT and return User DTO Json.

I could put PDF user list into context, but I had no time.

## How to run

You can run this code using docker.

1. Build

`docker build -t pdflas:latest .`

2. Run

`docker run -it -p 80:9090 pdflas:latest`

## How to auth

1. Go to localhost in your browser
2. Create a new account or go to Sign In and enter email: **test@t.pl**
3. Copy Access Code from the console. Due this reason We use flag -it during docker run.

`[ INFO ] 2025/03/26 16:37:01 (github.com/pecet3/las-test-pdf/cmd/router/auth.router.handleLogin:44)
↳ <Login> User with email: test@t.pl Access Code: X87F14`

I set up auth to does not send emails with Access Code (OTP), because it's not necessary in this case.
