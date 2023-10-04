start:
	docker compose -f "docker-compose.yml" up -d --build

stop:
	docker compose -f "docker-compose.yml" down

restart:
	docker compose -f "docker-compose.yml" down && docker compose -f "docker-compose.yml" up -d --build

migrate:
	cd sql/schema && goose postgres "postgresql://useradmin:qw123321@postgres/astro_db?sslmode=disable" up

