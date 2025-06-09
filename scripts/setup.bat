@echo off
echo 🚀 Setting up Reloop Backend...

echo 📦 Installing dependencies...
go mod tidy

echo 🗄️ Creating database directory...
if not exist "database" mkdir database

echo 🔄 Running migrations...
go run ./cmd/migrate

echo 🌱 Seeding database...
go run ./cmd/seed

echo ✅ Setup completed!
echo 🧪 You can now test the API:
echo    curl http://localhost:8080/v1/health
echo    go run ./cmd/api
pause