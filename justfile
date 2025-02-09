# Install Ent code-generation module
ent-install:
    go get entgo.io/ent/cmd/ent

# Generate Ent code
ent-gen:
    go generate ./ent

# Create a new Ent entity
ent-new name:
    go run entgo.io/ent/cmd/ent new {{name}}

# Run the application
run:
    clear
    go run cmd/web/main.go

# Run all tests
test:
    go test -count=1 -p 1 ./...

# Check for direct dependency updates
check-updates:
    go list -u -m all | grep "\[" | grep -v "indirect"
