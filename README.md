# payments
Simple example of a CRUD payments app utilising go-swagger, sqlx, and go-convey.

### Dependencies
- Go 
- postgresql server / client 
- dbmate (https://github.com/amacneil/dbmate) 

### Quick Start
  - Clone this git repo to your go path. 
   
    e.g. go/src/github/

   `git clone https://github.com/daesu/payments`

 - Set environment variables.

   `DATABASE_HOST=<postgres host>`
   
   `DATABASE_USERNAME=<postgres username>`
   
   `DATABASE_PASSWORD=<postgres password>`

   `DATABASE_NAME=<postgres db name>`
    e.g. DATABASE_NAME=payments

   `LOG_LEVEL` 
    Optional. Can be set to info, debug, error, fatal

   `DATABASE_URL=<postgres connection string>`
    e.g. postgresql://$DATABASE_USERNAME:$DATABASE_PASSWORD@$DATABASE_HOST/$DATABASE_NAME?sslmode=disable

 - Create & Seed database 
  `./db/migration.sh`

 - `make init` 

 - `make build`

 - `make run`
   Server will be started at localhost:8080 
   
### Tests
Tests use go-convey and are ran against a test database.

`make tests`

### Documentation
Documentation is autogenerated from swagger and available at `localhost:8080/v1/docs`
