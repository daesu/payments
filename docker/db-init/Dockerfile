# deploy alpine container
FROM alpine:latest

RUN apk --update add postgresql-client && rm -rf /var/cache/apk/*
RUN apk add --update bash && rm -rf /var/cache/apk/*
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache curl
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/download/v1.4.1/dbmate-linux-musl-amd64

RUN mkdir /code
WORKDIR /code

COPY . code/

RUN chmod +x /usr/local/bin/dbmate

# COPY env_file /app/env_file
# RUN source /app/env_file 

# RUN source ./code/code/env_file 
# Env variables 
# ENV DATABASE_HOST=db
# ENV DATABASE_PASSWORD=postgres
# ENV DATABASE_USERNAME=postgres
# ENV DATABASE_NAME=postgres
# ENV LOG_LEVEL=error
# ENV DATABASE_URL=postgresql://$DATABASE_USERNAME:$DATABASE_PASSWORD@$DATABASE_HOST/$DATABASE_NAME?sslmode=disable

RUN echo "Creating DB" 
CMD [ "./code/docker/db-init/wait-for-pg.sh", "db", "./code/docker/db-init/startup.sh" ]