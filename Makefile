APP = csv-normalization
DOCKER_IMAGE = $(APP)

# sets make build as default for all target
.PHONY: all
all: build

.PHONY: deps
deps:
	@go mod download
	@go mod tidy

.PHONY: run 
run: 
	@cd src && go run main.go INPUT=$(INPUT) OUTPUT=$(OUTPUT)

.PHONY: test
test: 
	@echo "Running unit tests..."
	@cd src/normalization && go test

.PHONY: build
docker-build:
	@echo "Building docker image..."
	@docker-compose build

.PHONY: run
docker-run: build
	@echo "Running dockerized app..."
	@docker-compose up

.PHONY: clean
clean: 
	@docker-compose down
    
#TODO: Add linting 