SERVICE = payments

generate: swagger

clean:
	rm -rf ./gen ./bin

swagger: clean
	mkdir gen
	swagger -q generate server -t gen -f swagger/payment.yaml

run:
	go run main.go

init:
	dep init 

build:
	make swagger
	dep ensure
	env GOOS=linux go build -o bin/$(SERVICE)

tests:
	make build 

	# Sets up test database
	./db/migration_testdb.sh

	# starts app and runs tests against testdb
	./test.sh