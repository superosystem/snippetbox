run-dev:
	go run main.go

setup-dev:
	docker compose -f docker/docker-compose-dev.yml up 