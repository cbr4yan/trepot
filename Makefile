.PHONY: deps build dev

deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod verify
	@echo "Done."

build: deps
	go build -o ./build/backend ./cmd/main.go

dev:
	docker-compose up
