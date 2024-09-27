include .env
export $(shell sed 's/=.*//' .env)

MIGRATE=migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL)
MIGRATE_CREATE=migrate create -ext sql -dir $(MIGRATION_PATH) -seq

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-create:
	$(MIGRATE_CREATE) $(name)
