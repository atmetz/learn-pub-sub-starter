package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	const connectionString = "amqp://guest:guest@localhost:5672/"

	connection, err := amqp.Dial(connectionString)
	if err != nil {
		fmt.Printf("Error connect to amqp: %v\n", err)
	}
	defer connection.Close()

	fmt.Println("Connection was successful")

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("exiting")
}
