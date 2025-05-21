create_migration:
	# Create a new migration
	@echo "Creating a new migration file..."
	migrate create -ext sql -dir db/migrations -seq $(name)

postgres_up:
	# Start the postgres container
	docker-compose up -d

postgres_down:
	# Stop the postgres container
	docker-compose down

create_db:
	# Create the database
	@echo "Creating the database..."
	docker exec -it fingreat_postgres createdb --username=fingreat_user --owner=fingreat_user fingreat_db
	docker exec -it fingreat_postgres_live createdb --username=fingreat_user --owner=fingreat_user fingreat_db_live

delete_db:
	# Delete the database
	@echo "Deleting the database..."
	docker exec -it fingreat_postgres dropdb --username=fingreat_user fingreat_db
	docker exec -it fingreat_postgres_live dropdb --username= fingreat_db_live

migrate_up:
	# Migrate up
	migrate -path db/migrations/ -database "postgres://fingreat_user:secret@localhost:5432/fingreat_db?sslmode=disable"  up
	migrate -path db/migrations/ -database "postgres://fingreat_user:secret@localhost:5432/fingreat_db_live?sslmode=disable"  up

migrate_down:
	# Migrate down
	migrate -path db/migrations/ -database "postgres://fingreat_user:secret@localhost:5432/fingreat_db?sslmode=disable"  down
	migrate -path db/migrations/ -database "postgres://fingreat_user:secret@localhost:5432/fingreat_db_live?sslmode=disable" down

sqlc:
	# sqlc generate
	sqlc generate

start:
	# Start the server
	@echo "Starting the server..."
	CompileDaemon -command="./fingreat_backend"

test:
	# Run the tests
	@echo "Running the tests..."
	go test -v -cover ./...
