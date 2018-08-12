package main

import (
	"os"
	"payment-sample/backend-api/infrastructure"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
