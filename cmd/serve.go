package cmd

import (
	"github.com/daesu/payments/gen/restapi"
	"github.com/daesu/payments/gen/restapi/operations"
)

func Start(api *operations.PaymentsAPI, portFlag int) error {
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = portFlag

	return server.Serve()
}
