# coach

## Setup
brew install golang-migrate

## Database
### Create Database
create database coach_development;
CREATE USER coach_user WITH ENCRYPTED PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE coach TO coach_user;
GRANT USAGE ON SCHEMA public TO coach_user;
RANT CREATE ON SCHEMA public TO coach_user;

### Create Migration
migrate create -ext sql -dir db/migrations -seq create_users_table
make migrate-create name=create_users_table

### Run Migration
make migrate-up


## Test
go test -v -cover ./...