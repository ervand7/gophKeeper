#! /bin/bash

testDBAddr='user=ervand password=ervand dbname=goph_keeper_test host=localhost port=5432 sslmode=disable'

kill_server() {
  lsof -t -i tcp:8080 | xargs kill -9
}

kill_server
DATABASE_DSN="$testDBAddr" ../server/cmd/main &

# shellcheck disable=SC1007
cd .. && DATABASE_DSN="$testDBAddr" go test ./... -coverprofile=coverage.out \
  -count 1 -v -p 1 && go tool cover -func coverage.out | grep total:
kill_server
