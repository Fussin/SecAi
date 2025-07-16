.PHONY: help build run test clean

help:
	@echo "AUTONOMOUSPEN AI - Makefile Commands"
	@echo "  make build - Build all containers"
	@echo "  make run   - Run the complete system"
	@echo "  make test  - Run all tests"
	@echo "  make clean - Clean build artifacts"

build:
	docker-compose build

run:
	docker-compose up -d

test:
	go test ./...
	python -m pytest python/

clean:
	docker-compose down -v
	rm -rf bin/
