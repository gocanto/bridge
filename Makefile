.PHONY: gocanto\:bridge

include .env

DB_NETWORK = bridge_network
APP_PATH = $(shell pwd)

app\:build:
	make docker:flush
	go mod tidy && \
	docker compose up app --build -d

app\:ssh:
	make docker:flush && \
	docker compose up app --build -d && \
	docker exec -it bridge_app /bin/bash

docker\:flush:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f
	rm -rf ./database/data

docker\:stop:
	docker compose down --volumes

docker\:status:
	docker compose ps


# -- Tests

test\:stripe:
	#https://docs.stripe.com/stripe-cli
	stripe trigger payment_intent.succeeded

tests:
	curl http://localhost:8080/service1 && \
	curl http://localhost:8080/service2 && \
	curl -H "X-Sender-ID: sender2" http://localhost:8080

# -- Local

# Requires the air package to be installed globally.
# https://github.com/cosmtrek/air
kill:
	#kill all air processes.
	lsof -t -i tcp:8080 | xargs kill -9
	lsof -t -i tcp:8081 | xargs kill -9

air:
	make kill
	air
