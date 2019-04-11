#!/bin/bash

export DATABASE_URL=postgresql://postgres:postgres@localhost/payment_test?sslmode=disable

# start the app in background and grab its pid
./bin/payments &
app=$!

sleep 1

# start tests
cd ./test && go test -v

sleep 1

# finish
kill $app
