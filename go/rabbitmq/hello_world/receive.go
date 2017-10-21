package main

import (
  "log"

  "github.com/streadway/amqp"
)

func main() {
	// connect to rmq server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	
	// create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	
	// declare the queue
	q, err := ch.QueueDeclare(
	  	"hello", // name
	  	false,   // durable
	  	false,   // delete when usused
	  	false,   // exclusive
	  	false,   // no-wait
	  	nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// consume a message from a queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	  
	forever := make(chan bool)
	  
	go func() {
		for d := range msgs {
		  log.Printf("Received a message: %s", d.Body)
		}
	}()
	  
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
