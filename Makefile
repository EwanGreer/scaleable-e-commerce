dev:
	docker compose up --build -d

user:
	docker compose up user --build -d

unit-test:
	go test -v ./...

generate-sql:
	sqlc -f ./services/user/repo/sqlc.yaml generate

test:
	go test ./... -v
