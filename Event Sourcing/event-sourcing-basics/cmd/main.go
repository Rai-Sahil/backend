package main

import (
	"log"

	"github.com/Rai-Sahil/backend/service"
	"github.com/Rai-Sahil/backend/store"
	"github.com/Rai-Sahil/backend/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "postgres://user:pass@localhost:5432/bank_events?sslmode=disable"
	eventStore, err := store.NewEventStore(dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	accountService := &service.AccountService{Store: *eventStore}
	handler := &transport.Handler{Account: accountService}

	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Fatal(r.Run(":8080"))
}
