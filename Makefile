.PHONY: build
build: ## Build application
	@go build main.go

.PHONY: clean
clean: ## Clean executable file
	@rm -f main

.PHONY: help
help: ## Help
	@./main -h

.PHONY: run
run: ## Run application
	@./main -h $(host) -p $(port) -t $(threads)