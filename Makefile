.PHONY: up down migrate-up migrate-down

up:
	docker compose up -d --build

down:
	docker compose down

migrate-up:
	migrate -path db/migrations -database "postgresql://aller:alleristhebest@localhost:5432/car_tool?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgresql://aller:alleristhebest@localhost:5432/car_tool?sslmode=disable" down