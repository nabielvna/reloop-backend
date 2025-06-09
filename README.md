# 🔄 Reloop Backend API

Marketplace backend system built with Go, PostgreSQL, and GORM.

## 🔧 Prerequisites

- **Go 1.20+** installed
- **PostgreSQL 15+** installed and running
- **pgAdmin4** (optional, for database management)

## 🚀 Quick Setup

### 1. Clone & Dependencies
```bash
git clone <repository-url>
cd reloop-backend
go mod tidy
go mod download
```

### 2. Database Setup
```bash
# Create database (via psql)
psql -U postgres -c "CREATE DATABASE reloop;"

# Or create via pgAdmin4:
# pgAdmin4 → Databases → Create → Database → Name: "reloop"
```

### 3. Environment Variables
```bash
# Application uses fallback environment variables:
# DB_ADDR=postgres://postgres:rafa2005@localhost:5432/reloop?sslmode=disable
# PORT=8080
# JWT_SECRET=your-secret-key
```

### 4. Database Migration & Seeding
```bash
# Run migration and populate sample data
go run ./cmd/migrate
```

### 5. Start Server
```bash
# Start the API server
go run ./cmd/api

# Expected output:
# ✅ Koneksi ke database berhasil
# 🚀 Server starting on :8080
```

## 🧪 API Testing Commands

### Health Check
```bash
curl http://localhost:8080/v1/health
# Expected: OK
```

## 🔐 Authentication Endpoints

### User Registration
```bash
curl -X POST http://localhost:8080/v1/auth/register \
-H "Content-Type: application/json" \
-d '{
  "email": "test@example.com",
  "password": "12345678",
  "userName": "testuser",
  "firstName": "Test",
  "lastName": "User",
  "phone": "081234567890"
}'
```

### User Login
```bash
curl -X POST http://localhost:8080/v1/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email": "seller@reloop.com",
  "password": "12345678"
}'

# Save the token from response for authenticated requests
```

### Admin Login
```bash
curl -X POST http://localhost:8080/v1/auth/admin/login \
-H "Content-Type: application/json" \
-d '{
  "email": "admin@reloop.com",
  "password": "12345678"
}'
```

## 📦 Items Endpoints

### Browse Items (Public)
```bash
# Get all items
curl http://localhost:8080/v1/items

# Get items by category
curl http://localhost:8080/v1/items?category_id=1

# Search items
curl http://localhost:8080/v1/items?search=macbook

# Get item by ID
curl http://localhost:8080/v1/items/1
```

### Create Item (Authenticated)
```bash
curl -X POST http://localhost:8080/v1/items \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_JWT_TOKEN" \
-d '{
  "name": "Test Product",
  "description": "Test product description",
  "price": 100000,
  "quantity": 10,
  "categoryId": 1
}'
```

### Update Item (Authenticated)
```bash
curl -X PUT http://localhost:8080/v1/items/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_JWT_TOKEN" \
-d '{
  "name": "Updated Product Name",
  "price": 150000,
  "quantity": 15
}'
```

### Delete Item (Authenticated)
```bash
curl -X DELETE http://localhost:8080/v1/items/1 \
-H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 👤 User Endpoints

### Get User Profile (Authenticated)
```bash
curl http://localhost:8080/v1/users/profile \
-H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Update User Profile (Authenticated)
```bash
curl -X PUT http://localhost:8080/v1/users/profile \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_JWT_TOKEN" \
-d '{
  "firstName": "Updated",
  "lastName": "Name",
  "phone": "081234567890"
}'
```

## 🏪 Seller Endpoints

### Register as Seller (Authenticated)
```bash
curl -X POST http://localhost:8080/v1/sellers/register \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_JWT_TOKEN" \
-d '{
  "businessName": "My Store",
  "businessAddress": "Jl. Example No. 123, Jakarta",
  "businessPhone": "021-12345678",
  "businessDescription": "Store description"
}'
```

### Get Seller Profile (Authenticated Seller)
```bash
curl http://localhost:8080/v1/sellers/profile \
-H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Get Seller Items (Authenticated Seller)
```bash
curl http://localhost:8080/v1/sellers/items \
-H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🏷️ Categories Endpoints

### Get All Categories (Public)
```bash
curl http://localhost:8080/v1/categories
```

### Create Category (Admin)
```bash
curl -X POST http://localhost:8080/v1/admin/categories \
-H "Content-Type: application/json" \
-H "Authorization: Bearer ADMIN_JWT_TOKEN" \
-d '{
  "name": "New Category",
  "description": "Category description"
}'
```

## 👨‍💼 Admin Endpoints

### Get All Users (Admin)
```bash
curl http://localhost:8080/v1/admin/users \
-H "Authorization: Bearer ADMIN_JWT_TOKEN"
```

### Get All Items (Admin)
```bash
curl http://localhost:8080/v1/admin/items \
-H "Authorization: Bearer ADMIN_JWT_TOKEN"
```

### Approve Item (Admin)
```bash
curl -X PUT http://localhost:8080/v1/admin/items/1/approve \
-H "Authorization: Bearer ADMIN_JWT_TOKEN"
```

### Reject Item (Admin)
```bash
curl -X PUT http://localhost:8080/v1/admin/items/1/reject \
-H "Content-Type: application/json" \
-H "Authorization: Bearer ADMIN_JWT_TOKEN" \
-d '{
  "reason": "Does not meet quality standards"
}'
```

## 🔍 Sample Test Credentials

### Regular Users
```
Email: seller@reloop.com
Password: 12345678
Role: user (can become seller)

Email: user@reloop.com  
Password: 12345678
Role: user
```

### Admin
```
Email: admin@reloop.com
Password: 12345678
Role: admin
```

## 📊 Database Tables

- `users` - User accounts
- `admins` - Admin accounts  
- `sellers` - Seller profiles
- `categories` - Product categories
- `items` - Products/items
- `product_reviews` - Item reviews
- `fraud_reports` - Fraud reports

## 🛠️ Development Commands

```bash
# Build binary
go build -o ./bin/main.exe ./cmd/api/

# Run tests
go test ./...

# Format code
go fmt ./...

# Clean module cache
go clean -modcache

# Rebuild dependencies
go mod tidy && go mod download
```

## 📝 Common Issues

### Server exits silently
- Check Go version compatibility
- Verify PostgreSQL service is running
- Check database connection credentials

### Database connection failed
- Verify PostgreSQL service: `Get-Service postgresql*`
- Check credentials: username=postgres, password=rafa2005
- Ensure database 'reloop' exists

### Module download errors
- Clear cache: `go clean -modcache`
- Reset proxy: `go env -w GOPROXY=https://proxy.golang.org,direct`
- Rebuild: `go mod tidy`