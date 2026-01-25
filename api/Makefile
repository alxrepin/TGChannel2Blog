setup: local
	go mod tidy
	build
	swagger

#create_migration:
#	migrate create -ext=sql -dir=migrations -seq "$$DATABASE_URL"

#migrate_down:
#	migrate -path=./migrations -database "$$DATABASE_URL" -verbose down

build:
	go build -o ./bin/server cmd/server/main.go \
		&& go build -o ./bin/sync cmd/sync/main.go \
		&& go build -o ./bin/subscriber cmd/subscriber/main.go \
		&& go build -o ./bin/media_processor cmd/media_processor/main.go

local:
	docker compose --env-file .env -f ./.docker/compose/compose.yml up -d $(if $(filter 1, $(BUILD)),--build,)

stop:
	@docker ps --filter "name=tgchannel2blog" --format "{{.ID}}" | xargs -r docker stop

test:
	go test ./...

swagger:
	swag init -g cmd/server/main.go

.PHONY: swagger test stop local build setup
