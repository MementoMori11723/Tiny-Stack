.DEFAULT_GOAL := build

help:
	@echo "Usage:"
	@echo "  make build   - Generate code and move it to the src directory"
	@echo "  make run     - Run the Go application"
	@echo "  make clean   - Remove generated templ files"
	@echo "  make dev     - Run the development server and then the app"
build:
	templ generate && mv ./src/pages/*.go ./src
run:
	go run ./src/*.go
clean:
	find ./src -type f -name '*templ.go' -exec rm {} \;
dev: ./dev/dev.go
	go run ./dev/dev.go && make run