# Go builder container
FROM golang:1.11 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/daesu/payments
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /payments .

# deploy alpine container
FROM alpine:latest

RUN apk --update add postgresql-client && rm -rf /var/cache/apk/*
RUN apk add --update bash && rm -rf /var/cache/apk/*
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache curl
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/download/v1.4.1/dbmate-linux-musl-amd64

RUN mkdir /app
WORKDIR /app

COPY --from=builder /payments /app

COPY db /app/db/

RUN chmod +x /app/db/migration.sh
RUN chmod +x /usr/local/bin/dbmate

# Env variables 
ENV DATABASE_HOST=db
ENV DATABASE_PASSWORD=postgres
ENV DATABASE_USERNAME=postgres
ENV DATABASE_NAME=postgres
ENV LOG_LEVEL=error
ENV DATABASE_URL=postgresql://$DATABASE_USERNAME:$DATABASE_PASSWORD@$DATABASE_HOST/$DATABASE_NAME?sslmode=disable

EXPOSE 8080
CMD ["/app/payments"]
