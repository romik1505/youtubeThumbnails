CURRENT_DIR = $(shell pwd)
LOCAL_BIN=$(CURRENT_DIR)/bin
VALUES=$(CURRENT_DIR)/.o3/k8s/values_local.yaml

ifndef SQLITE
$(eval SQLITE=$(shell cat $(VALUES) | grep -i "sqlite" -A1 | sed -n '2p;2q' | sed -e 's/[ \t]*value://g'))
SQLITE:=$(addprefix $(CURRENT_DIR)/, $(SQLITE))
endif

bin-depth:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.5.3
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.1

db\:create:
	$(LOCAL_BIN)/goose -dir migrations create "$(NAME)" sql

db\:up:
	@mkdir -p db
	$(LOCAL_BIN)/goose -dir migrations sqlite3 "$(SQLITE)" up

db\:down:
	$(LOCAL_BIN)/goose -dir migrations sqlite3 "$(SQLITE)" down

run:
	@go run ./cmd/server/main.go

generate:
	protoc -I ./api --go_out ./pkg/api/ --go_opt paths=source_relative \
	--go-grpc_out ./pkg/api/ --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./pkg/api/ --grpc-gateway_opt paths=source_relative \
	--openapiv2_out ./pkg/api/ --openapiv2_opt logtostderr=true \
	./api/thumbnails/thumbnails.proto

mocks:
	mockgen -source=./internal/app/service/service.go -destination=./internal/pkg/mock/service/mock_service.go
	mockgen -source=./internal/app/store/thumbnail/thumbnail.go -destination=./internal/pkg/mock/repository/mock_rep.go

test:
	SQLITE=$(SQLITE) VALUES=$(VALUES) go test -v ./...

lint:
	$(LOCAL_BIN)/golangci-lint run ./...
