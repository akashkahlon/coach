
# Coach

**Coach** is a developer coaching platform that tracks and improves code quality, efficiency, and performance through metrics and daily/weekly learning prompts. This repository contains the backend services for managing organizations, developers, and GitHub data integration.

## Table of Contents

- [Coach](#coach)
	- [Table of Contents](#table-of-contents)
	- [Setup](#setup)
	- [Database](#database)
		- [Create Database](#create-database)
		- [Create Migration](#create-migration)
		- [Run Migration](#run-migration)
	- [Testing](#testing)

## Setup

To get started, ensure you have the following prerequisites installed:

1. **Golang Migrate**: A migration tool for Go.
   ```bash
   brew install golang-migrate
   ```

2. **Make**: Ensure `make` is installed on your system for managing common tasks.

## Database

### Create Database

Run the following SQL commands to set up your PostgreSQL database:

```sql
CREATE DATABASE coach_development;
CREATE USER coach_user WITH ENCRYPTED PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE coach_development TO coach_user;
GRANT USAGE ON SCHEMA public TO coach_user;
GRANT CREATE ON SCHEMA public TO coach_user;
```

### Create Migration

To create a new migration, use the following commands:

```bash
migrate create -ext sql -dir db/migrations -seq create_users_table
```

Alternatively, if you have `make` set up, you can use:

```bash
make migrate-create name=create_users_table
```

### Run Migration

Run all database migrations using:

```bash
make migrate-up
```

## Testing

Run the full test suite with code coverage:

```bash
go test -v -cover ./...
```
