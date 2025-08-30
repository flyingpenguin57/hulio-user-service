# User Service

This service provides a set of user-related functionalities that can be reused across different applications. It handles common user operations, allowing you to focus on other parts of your projects without reimplementing user management.

## Features

The User Service currently supports the following functions:

- **Login**: Authenticate users with username/password.
- **Register**: Create a new user account.
- **Get User Info**: Retrieve details of a specific user.
- **Edit User Info**: Update user profile information.
- **Delete User**: Remove a user from the system.

## Usage

You can integrate this service into your applications to manage user-related operations easily. Simply call the corresponding API endpoints or service methods as needed.

## Benefits

- **Reusability**: Use the same user service across multiple projects.
- **Consistency**: Ensure uniform handling of user data and authentication.
- **Maintainability**: Centralized user management makes updates and bug fixes easier.

## How to use
### Init database
you can run mysql script in scripts to init database.

### Environment variables

Set the following environment variables before running the service:

- **ENV**: Specify enviroment, PROD for production and TEST for running test cases.

- **MYSQL_DSN_PROD**: MySQL DSN used when `ENV=PROD`.
  - Example:
    ```bash
    export MYSQL_DSN_PROD="user:password@tcp(127.0.0.1:3306)/hulio_user?charset=utf8mb4&parseTime=True&loc=Local"
    ```

- **MYSQL_DSN_TEST**: MySQL DSN used when `ENV=TEST`.
  - Example:
    ```bash
    export MYSQL_DSN_TEST="user:password@tcp(127.0.0.1:3306)/hulio_user_test?charset=utf8mb4&parseTime=True&loc=Local"
    ```

- **JWT_PRIVATE_KEY**: RSA private key (PEM format) used to sign JWTs (RS256). Required for login/token generation.
  - You can generate a key pair via the helper script, then load the private key into the env variable:
    ```bash
    # generate RSA key pair
    go run ./scripts/key/gen_key.go

    # load private key (adjust path if needed)
    export JWT_PRIVATE_KEY="$(cat ./scripts/private_key.pem)"
    ```

- **JWT_PRIVATE_KEY_TEST**: Same as JWT_PRIVATE_KEY, which is used for running test cases.

### Run application
please inject enviromental variables [JWT_PRIVATE_KEY, MYSQL_DSN_PROD] before run application
、、、bash
ENV=prod go run main.go
```