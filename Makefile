dev:
	docker compose up

user:
	docker compose up user --build

unit-test:
	go test -v ./...
