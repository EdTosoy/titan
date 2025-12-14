DSN = postgres://localhost/titan?sslmode=disable


.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


.PHONY: run
run:
	go run ./cmd/api -db-dsn=${DSN}


.PHONY: db/migrations/up
db/migrations/up:
	migrate -path=./migrations -database=${DSN} up


.PHONY: db/migrations/down
db/migrations/down:
	migrate -path=./migrations -database=${DSN} down 1


.PHONY: db/migrations/new
db/migrations/new:
	migrate create -seq -ext=.sql -dir=./migrations ${name}


