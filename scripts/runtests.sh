#! /bin/bash

testDBAddr='user=ervand password=ervand dbname=goph_keeper_test host=localhost port=5432 sslmode=disable'

kill_server() {
  lsof -t -i tcp:8080 | xargs kill -9
}

kill_server
../server/cmd/main -d "$testDBAddr" &
DATABASE_DSN="$testDBAddr" go test ../client/internal/requests \
  -count 1 -v -p 10 -bench=. -cpu 8 -benchmem
kill_server

DATABASE_DSN="$testDBAddr" go test \
  ../client/internal/config \
  ../pkg/algorithms \
  ../pkg/cert \
  ../pkg/encryption \
  ../server/internal/config \
  ../server/internal/storage \
  -count 1 -v -p 10 -bench=. -cpu 8 -benchmem
