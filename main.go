package main

import (
	"flag"
	"os"

	"github.com/daesu/payments/gen/restapi"
	"github.com/daesu/payments/gen/restapi/operations"
	"github.com/daesu/payments/health"
	"github.com/daesu/payments/payment"

	"github.com/daesu/payments/cmd"
	log "github.com/daesu/payments/logging"
	"github.com/go-openapi/loads"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// initDB
func initDB() (*sqlx.DB, error) {
	log.Println("Initializing DB")

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("err", err)
	}
	db.SetMaxOpenConns(2)

	return db, nil
}

func start(pDB *sqlx.DB) {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal("unable to get Hostname", err)
	}
	log.WithFields(logrus.Fields{
		"Host": host,
	}).Info("Service Startup")

	var portFlag = flag.Int("port", 8080, "Port to listen for web requests on")

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal("Invalid swagger file for initializing", err)
	}

	api := operations.NewPaymentsAPI(swaggerSpec)

	// Health setup
	healthService := health.New()
	health.Configure(api, healthService)

	// Payments package endpoints
	paymentRepo := payment.NewRepository(pDB)
	paymentService := payment.NewService(paymentRepo)
	payment.Configure(api, paymentService)

	if err := cmd.Start(api, *portFlag); err != nil {
		log.Fatal("Failed to start", err)
	}
}

func main() {

	// DB setup
	pDB, err := initDB()
	if err != nil {
		log.Fatal("couldn't connect to database", err)
	}

	start(pDB)
}
