#!/bin/bash

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/$GOPATH/bin
set -e
set -u

export TEST_DB_NAME=payment_test
export DATABASE_URL=postgresql://$DATABASE_USERNAME:$DATABASE_PASSWORD@$DATABASE_HOST/$TEST_DB_NAME?sslmode=disable

echo "Migration started"

# echo "Deleting database..."
dbmate drop

sleep 2

echo "Creating database..."
dbmate up

sleep 3

echo "Migration finished"

echo "Seeding started"
echo $DATABASE_HOST

export PGPASSWORD=$DATABASE_PASSWORD

psql \
    -X \
    -U $DATABASE_USERNAME \
    -h $DATABASE_HOST \
    -w \
    -a \
    -f ./db/seed.sql \
    --echo-all \
    --single-transaction \
    --set AUTOCOMMIT=off \
    --set ON_ERROR_STOP=on \
    $TEST_DB_NAME

echo "seed script successful"