API_SPEC=api/openapi.yaml
API_GEN_DIR=internal\shared\openapi
API_GEN_FILE=$(API_GEN_DIR)\openapi.gen.go
BINARY_NAME=academy

generate-openapi:
	if not exist $(API_GEN_DIR) mkdir $(API_GEN_DIR)
	oapi-codegen -generate types,server -package openapi -o $(API_GEN_FILE) $(API_SPEC)

run:
	go run cmd/academy/main.go

build:
	if not exist bin mkdir bin
	go build -o bin\$(BINARY_NAME).exe cmd/academy/main.go

test:
	go test ./... -v

clean:
	if exist $(API_GEN_FILE) del $(API_GEN_FILE)
	if exist bin rmdir /s /q bin
