.PHONY: tea-skeleton/server/grpc
tea-skeleton/server/grpc:
	cd "./app/tea-skeleton"; \
	go build -o "./cmd/server/grpc/main" "./cmd/server/grpc/main.go"; \
	"./cmd/server/grpc/main"

.PHONY: tea-skeleton/server/http
tea-skeleton/server/http:
	cd "./app/tea-skeleton"; \
	go build -o "./cmd/server/http/main" "./cmd/server/http/main.go"; \
	"./cmd/server/http/main"
