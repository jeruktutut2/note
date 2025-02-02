package main

import (
	"context"
	"net/http"
	"note-golang-kafka/consumers"
	"note-golang-kafka/routes"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// config := &kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092",
	// 	"group.id":          "golang",
	// 	"auto.offset.reset": "earliest",
	// }

	e := echo.New()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	consumer := consumers.NewKafkaConsumer(ctx)
	defer consumer.Close()

	routes.SetKafkaRoute(e)

	// config := &kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092",
	// }
	// producer, err := kafka.NewProducer(config)
	// if err != nil {
	// 	log.Fatalln("error creating producer:", err)
	// }
	// kafkaService := services.NewKafkaService(producer)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
