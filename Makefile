.PHONY: setup migrate seed clean run test

# Setup everything
setup:
	@echo "🚀 Setting up Reloop Backend..."
	go mod tidy
	mkdir -p database
	make migrate
	make seed
	@echo "✅ Setup completed!"

# Run migrations
migrate:
	@echo "🔄 Running migrations..."
	go run ./cmd/migrate

# Seed database
seed:
	@echo "🌱 Seeding database..."
	go run ./cmd/seed

# Clean database
clean:
	@echo "🧹 Cleaning database..."
	rm -rf database/*.db
	@echo "✅ Database cleaned!"

# Run application
run:
	@echo "🚀 Starting server..."
	go run ./cmd/api

# Build application
build:
	@echo "🔨 Building application..."
	go build -o ./bin/main.exe ./cmd/api/

# Test API endpoints
test:
	@echo "🧪 Testing API..."
	curl http://localhost:8080/v1/health
	@echo "🔍 Browse items:"
	curl http://localhost:8080/v1/items

# Reset everything
reset: clean migrate seed
	@echo "🔄 Database reset completed!"

# Development workflow
dev: setup run