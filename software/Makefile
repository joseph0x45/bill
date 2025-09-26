.PHONY: up down logs psql reset migrate-up migrate-down migrate-new

MIGRATE_DB_URL=bill.db
MIGRATE_PATH=./migrations
MIGRATE_BIN=gomigrate

build:
	go build .

build_release:
	go build -tags release .

migrate-up:
	$(MIGRATE_BIN) -path $(MIGRATE_PATH) -database $(MIGRATE_DB_URL) up

migrate-down:
	$(MIGRATE_BIN) -path $(MIGRATE_PATH) -database $(MIGRATE_DB_URL) down 1

migrate-new:
	@test -n "$(name)" || (echo ">> name= required"; exit 1)
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATE_PATH) -seq $(name)

