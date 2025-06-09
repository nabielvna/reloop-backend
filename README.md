# Reloop Backend

Clean Architecture Go backend with JWT authentication and CRUD operations.

## 🚀 Quick Start

### Windows:
```bash
# Run setup script
scripts\setup.bat

# Or manually:
make setup


# Clean cache
go mod tidy

# Test build (1)
go build -o ./bin/main.exe ./cmd/api/

# Test run (2)
./bin/main.exe


# DEVELOPMENT COMMANDS : 

# Setup everything (first time)
make setup

# Run migrations only
make migrate

# Seed database only  
make seed

# Clean database
make clean

# Reset database (clean + migrate + seed)
make reset

# Run application
make run

# Build application
make build

# Test API
make test


