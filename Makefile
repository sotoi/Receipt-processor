.PHONY: fmt
fmt:
	@echo "Formatting code.."
	go fmt ./...

.PHONY: build
build:
	@echo "Building the application.."
	go build -o bin/main main.go
.PHONY: run
run: build
	@echo "Running the application.."
	./bin/main