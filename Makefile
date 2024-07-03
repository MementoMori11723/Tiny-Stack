.DEFAULT_GOAL := build

help:
	@echo "Usage:"
	@echo "  make         - Generate code and move it to the src directory"
	@echo "  make run     - Run the Go application"
	@echo "  make pkg     - Build the Go application and move it to the bin directory"
	@echo "  make help    - Show this help message"
build:
	templ generate -path ./pages/templates && mv ./pages/templates/*.go ./pages
run:
	go run main.go
pkg:
	go build -o server ./src && mkdir bin && mv server ./bin