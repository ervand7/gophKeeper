.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20

.PHONY: runserver # Run server with db in docker (docker-compose up)
runserver:
	docker-compose up

.PHONY: downserver # Down server with db (docker-compose down)
downserver:
	docker-compose down

.PHONY: runclient-darwin-amd64 # Run client for darwin-amd64 with "qwerty" encryptionKey
runclient-darwin-amd64:
	cd client/cmd && encryptionKey=qwerty ./client_darwin_amd64

.PHONY: runclient-darwin-arm64 # Run client for darwin-arm64 with "qwerty" encryptionKey
runclient-darwin-arm64:
	cd client/cmd && encryptionKey=qwerty ./client_darwin_arm64

.PHONY: runclient-linux-amd64 # Run client for linux-amd64 with "qwerty" encryptionKey
runclient-linux-amd64:
	cd client/cmd && encryptionKey=qwerty ./client_linux_amd64

.PHONY: runclient-windows-amd64 # Run client for windows-amd64 with "qwerty" encryptionKey
runclient-windows-amd64:
	cd client/cmd && encryptionKey=qwerty ./client_windows_amd64

.PHONY: build-server # Build server executable file
build-server:
	cd server/cmd && go build main.go

.PHONY: build-clients # Build client executable files for darwin-arm64, darwin-amd64, windows-amd64 and linux-amd64
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

.PHONY: proto # Generate proto-files
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/grpc.proto

.PHONY: goose-down # Down to initial migration
goose-down:
	cd migrations && export GOOSE_DRIVER=postgres \
		&& export GOOSE_DBSTRING='host=localhost user=ervand password=ervand database=goph_keeper' && goose down

.PHONY: runtests # Run unit-tests
runtests:
	cd scripts/ && ./runtests.sh

.PHONY: check-coverage # Check tests coverage
check-coverage:
	cd scripts/ && ./check_tests_coverage.sh
