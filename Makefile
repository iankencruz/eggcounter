.PHONY: all build run

all: build run

build:
	@echo "Building the backend..."
	cd backend && go build -o tmp/main ./cmd/api/

run:
	@echo "Running the backend..."
	cd backend && ./tmp/main


sv-install:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

sv-dev:
	@echo "Installing frontend dependencies and starting development server..."
	cd frontend && npm run dev

sv-build:
	@echo "Installing frontend dependencies and building for production..."
	cd frontend && npm run build
