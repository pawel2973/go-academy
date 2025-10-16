# Makefile for Go Academy project

API_SPEC=api/openapi.yaml
API_GEN_DIR=internal/shared/openapi
API_GEN_FILE=$(API_GEN_DIR)/openapi.gen.go

generate-openapi:
	@echo "Generating Go types and server interfaces from OpenAPI specification..."
	oapi-codegen -generate types,server -package openapi -o $(API_GEN_FILE) $(API_SPEC)
	@echo "OpenAPI code generated successfully."

run:
	@echo "Running Go Academy application..."
	go run cmd/academy/main.go

test:
	@echo "Running tests..."
	go test ./... -v

clean:
	@echo "Cleaning generated files..."
	rm -f $(API_GEN_FILE)
