#!/bin/bash
VERSION=$(shell cat VERSION)

run:
	go run main.go

docker-build:
	docker build --tag ecojuntak/go-rest-api:${VERSION} --tag ecojuntak/go-rest-api:latest .

docker-push:
	docker push ecojuntak/go-rest-api:${VERSION} 
	docker push ecojuntak/go-rest-api:latest