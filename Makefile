SERVICE = payments
DOCKER_IMAGE_TAG        ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
DOCKER_REPO             ?= daesu
DOCKER_IMAGE_NAME		?= payments

docker-run:
	docker-compose up

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

	# Set up test database
	./db/migration_testdb.sh

	# starts app and runs tests against testdb
	./test/test.sh