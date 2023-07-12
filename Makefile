#!/bin/bash
VERSION=$(shell cat VERSION)

build:
	go build -o bin/go-rest-api main.go

run: build
	./bin/go-rest-api serve

compose-recreate:
	docker compose down
	rm -rf postgres-data
	docker compose up -d


migrate:
	./bin/go-rest-api migrate

docker-build:
	docker build --tag ecojuntak/go-rest-api:${VERSION} --tag ecojuntak/go-rest-api:latest .

docker-push:
	docker push ecojuntak/go-rest-api:${VERSION} 
	docker push ecojuntak/go-rest-api:latest