# Clean cache
go mod tidy

# Test build
go build -o ./bin/main.exe ./cmd/api/

# Test run
./bin/main.exe


