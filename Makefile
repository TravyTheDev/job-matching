include .env
run:
	@go run .

up:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${DB_URL} goose -dir=db/migrations up

down:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${DB_URL} goose -dir=db/migrations down