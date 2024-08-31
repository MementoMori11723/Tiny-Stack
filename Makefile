default_goal: run
run:
	@echo "Running Server..."
	@go run .

generate:
	@echo "Generating code..."
	@templ generate

test:
	@echo "Running tests..."
	@go test -v ./...

