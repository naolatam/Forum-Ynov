## forum

### Objectives

This project consists in creating a web forum that allows :

- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.

#### MySQL

In order to store the data in your forum (like users, posts, comments, etc.) you will use the database library MySQL.

MySQL is a popular choice as an embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

To structure your database and to achieve better performance we highly advise you to take a look at the [entity relationship diagram](https://www.smartdraw.com/entity-relationship-diagram/) and build one based on your own database.

- You must use at least one SELECT, one CREATE and one INSERT queries.

To know more about MySQL you can check the [MySQL page](https://dev.mysql.com/doc/).

#### Authentication

In this segment the client must be able to `register` as a new user on the forum, by inputting their credentials. You also have to create a `login session` to access the forum and be able to add posts and comments.

You should use cookies to allow each user to have only one opened session. Each of this sessions must contain an expiration date. It is up to you to decide how long the cookie stays "alive". The use of UUID is a Bonus task.

Instructions for user registration:

- Must ask for email
  - When the email is already taken return an error response.
- Must ask for username
- Must ask for password
  - The password must be encrypted when stored (this is a Bonus task)

The forum must be able to check if the email provided is present in the database and if all credentials are correct. It will check if the password is the same with the one provided and, if the password is not the same, it will return an error response.

#### Communication

In order for users to communicate between each other, they will have to be able to create posts and comments.

- Only registered users will be able to create posts and comments.
- When registered users are creating a post they can associate one or more categories to it.
  - The implementation and choice of the categories is up to you.
- The posts and comments should be visible to all users (registered or not).
- Non-registered users will only be able to see posts and comments.

#### Likes and Dislikes

Only registered users will be able to like or dislike posts and comments.

The number of likes and dislikes should be visible by all users (registered or not).

#### Filter

You need to implement a filter mechanism, that will allow users to filter the displayed posts by :

- categories
- created posts
- liked posts

You can look at filtering by categories as subforums. A subforum is a section of an online forum dedicated to a specific topic.

Note that the last two are only available for registered users and must refer to the logged in user.

#### Docker

For the forum project you must use Docker.
> Your MySQL database must be deployed in a dedicated Docker container with proper volume persistence. The backend must connect to this container using Docker networking or environment variables.
 You can read about docker basics in the [ascii-art-web-dockerize](https://github.com/SofyaOspan/-ascii-art-web-dockerize/blob/main/README.md) subject.

### Instructions

- You must use **MySQL**.
> Avoid hardcoding database credentials. Use environment variables or configuration files so that your application remains portable and testable.

> Be particularly careful to handle database-specific errors (e.g. duplicate entries, constraint violations, connection timeouts).

> Ensure your MySQL container is not exposed externally (no `0.0.0.0:3306`), only accessible from the backend.

- You must handle website errors, HTTP status.
- You must handle all sort of technical errors.
- The code must respect the [**good practices**](https://ytrack.learn.ynov.com/git/root/public/src/branch/master/subjects/good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

### Allowed packages

- All [standard Go](https://golang.org/pkg/) packages are allowed.
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [gofrs/uuid](https://github.com/gofrs/uuid) or [google/uuid](https://github.com/google/uuid)

> You must not use any frontend libraries or frameworks like React, Angular, Vue etc.

This project will help you learn about:

- The basics of web :
  - HTML
  - HTTP
  - [Sessions](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html#session-management-waf-protections) and [cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies)
- Using and [setting up Docker](https://docs.docker.com/get-started/)
  - Containerizing an application
  - Compatibility/Dependency
  - Creating images
- SQL language
  - Manipulation of databases
- The basics of encryption
---

### Features

The following extensions of the forum are available as separate modules. Each one has its own README file:

- [Advanced Features](advanced-features/README.md): Notifications, user activity tracking, post editing/deletion.
- [Authentication](authentication/README.md): OAuth login using GitHub and Google.
- [Image Upload](image-upload/README.md): Support for image attachments in posts (JPEG, PNG, GIF).
- [Moderation](moderation/README.md): User roles and content approval system.
- [Security](security/README.md): HTTPS, rate limiting, session hardening and encryption.

Each feature must be implemented **on top of the base forum**, using consistent design and architecture.
