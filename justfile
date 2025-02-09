# Install Ent code-generation module
ent-install:
    go get entgo.io/ent/cmd/ent

# Generate Ent code
ent-gen:
    go generate ./ent

# Create a new Ent entity
ent-new name:
    go run entgo.io/ent/cmd/ent new {{name}}

# Install templ code-generation tool
templ-install:
    go install github.com/a-h/templ/cmd/templ@latest

# Generate templ code
templ-gen:
    templ generate

# Run the application
run:
    clear
    go run cmd/web/main.go

# Run all tests
test: templ-gen
    go test -count=1 -p 1 ./...

# Check for direct dependency updates
check-updates:
    go list -u -m all | grep "\[" | grep -v "indirect"
