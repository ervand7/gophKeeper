.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20

.PHONY: runbackground # Run server with db
runbackground:
	docker-compose up -d && \
	sleep 30 && \
	cd server/cmd && \
	DATABASE_DSN='user=ervand password=ervand dbname=goph_keeper host=localhost port=5466 sslmode=disable' ./main

.PHONY: build-server # Builds server executable file
build-server:
	cd server/cmd && go build main.go

.PHONY: build-clients # Builds client executable files for darwin-arm64, darwin-amd64, windows-amd64 and linux-amd64
build-clients:
	cd client/cmd && GOOS=darwin GOARCH=arm64 go build \
		-ldflags="-X 'main.buildVersion=v1.0.0' -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" \
		-o client_darwin_arm64 main.go && GOOS=darwin GOARCH=amd64 go build \
		-ldflags="-X 'main.buildVersion=v1.0.0' -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" \
		-o client_darwin_amd64 main.go && GOOS=windows GOARCH=amd64 go build \
		-ldflags="-X 'main.buildVersion=v1.0.0' -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" \
		-o client_windows_amd64 main.go && GOOS=linux GOARCH=amd64 go build \
		-ldflags="-X 'main.buildVersion=v1.0.0' -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" \
		-o client_linux_amd64 main.go

.PHONY: proto # Generates proto-files
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/grpc.proto

.PHONY: goose-down # Down to initial migration
goose-down:
	cd server/migrations && export GOOSE_DRIVER=postgres \
		&& export GOOSE_DBSTRING='host=localhost user=ervand password=ervand database=goph_keeper' && goose down

.PHONY: runserver # Runs server
runserver:
	cd server/cmd && DATABASE_DSN='user=ervand password=ervand dbname=goph_keeper host=localhost port=5466 sslmode=disable' ./main

.PHONY: runclient # Runs client_darwin_arm64
runclient:
	cd client/cmd && encryptionKey=qwerty ./client_darwin_arm64

.PHONY: runtests # Runs unit-tests
runtests:
	cd scripts/ && ./runtests.sh

.PHONY: check-coverage # Checks tests coverage
check-coverage:
	cd scripts/ && ./check_tests_coverage.sh
