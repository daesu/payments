#!/bin/bash

cd /code/code

echo "kill database..."
dbmate down

echo "create database..."
dbmate up

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
    $DATABASE_NAME

echo "seed script successful"