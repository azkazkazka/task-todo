# default values for var
DB_USER := postgres
DB_PASSWORD := password
DB_HOST := localhost
DB_PORT := 5432
DB_NAME := tasktodo
SSL_MODE := disable

migrate-up:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)" -verbose down

test:
	go test ./controllers
