default_goal: build

build:
	@echo "Building..."
	find . -name "*templ.go" | while read file; do 
		mv "$file" /path/to/destination/ 
	done
	@echo "Done."
