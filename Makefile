dev:
	docker compose up

users:
	docker compose up users --build

unit-test:
	go test -v ./...
