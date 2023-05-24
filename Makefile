runLocal:
	docker compose -f docker-compose.local.yml up -d --build

stopLocal:
	docker compose -f docker-compose.local.yml down

.PHONY: runLocal stopLocal
