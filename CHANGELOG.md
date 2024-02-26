## 1.0.0 (2024-02-26)

### Feat

- **auth-middleware**: auth middleware now passes along user information to controllers
- **controller**: refactor the service functions and add a reusable generic CRUD controller
- **service**: add service functions for querying post, category, and blog data from the database

### Refactor

- **app**: update naming convention to match the Golang standard

## 0.3.0 (2024-02-01)

### Feat

- **database**: add database tables for blogs, posts, and categoriess
- **user**: add a route for retrieving logged in user information
- **auth**: add a middleware for validating login status
- **session**: implement user session management functions
- **oauth**: integrate OAuth2 in the login flow

### Fix

- **user**: fix an issue where existing user cannot connect to their Google account

## 0.2.0 (2023-12-13)

### Feat

- **user**: implement user registration
- **project**: add gorm, app config, and basic database connection function

### Refactor

- **response**: refactor the error response functions to be reusable

## 0.1.0 (2023-12-10)

### Feat

- **project**: set up project structure and basic routing
- **project**: init
