.DEFAULT_GOAL := build

build:
	templ generate && mv ./src/pages/*.go ./src
run:
	go run ./src/*.go
clean:
	find ./src -type f -name '*templ.go' -exec rm {} \;