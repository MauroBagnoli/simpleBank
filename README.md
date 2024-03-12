# SimpleBank

SimpleBank is a lightweight, Go-based banking application designed to demonstrate core banking operations like account management, transfers, and transaction logging. It's built with a focus on clean architecture, best practices, and scalability.

## Features

- **Account Management**: Create, update, and manage bank accounts.
- **Transaction Handling**: Support for deposits, withdrawals, and account transfers.
- **Secure Authentication**: Implement JWT-based authentication to secure endpoints.
- **Database Migrations**: Integrated support for schema migrations using golang-migrate.
- **RESTful API**: Easy-to-use API for all banking operations.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.20 or later)
- Docker and Docker Compose (for running with containers)
- PostgreSQL (for the database)

### Installing

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/simpleBank.git
cd simpleBank
```

2. **Run database migrations**

Make sure your PostgreSQL database is running. Then, execute:

```bash
make migrateup
```
2. **Start the server**

```bash
make server
```
2. **Run the tests**

```bash
make test
```

## Built With

- [Go](https://golang.org/) - The Go programming language
- [PostgreSQL](https://www.postgresql.org/) - Open Source Relational Database
- [JWT-go](https://github.com/dgrijalva/jwt-go) - JSON Web Tokens for Go
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
