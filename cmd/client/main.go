package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")
	const connectionString = "amqp://guest:guest@localhost:5672/"

	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}
	defer connection.Close()

	fmt.Println("Peril game client connected to RabbitMQ!")

	userName, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatalf("Error getting user name: %v\n", err)
	}

	_, newQueue, err := pubsub.DeclareAndBind(
		connection,
		routing.ExchangePerilDirect,
		routing.PauseKey+"."+userName,
		routing.PauseKey,
		pubsub.Transient,
	)
	if err != nil {
		log.Fatalf("could not subscribe to pause: %v", err)
	}

	fmt.Printf("Queue %v declared and bound!\n", newQueue.Name)

	// wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("RabbitMQ connection closed.")

}
