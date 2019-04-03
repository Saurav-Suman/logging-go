package publisher

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func sendThisErrorOnPriority(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func publisher() {
	conn, err := amqp.Dial(os.Getenv("LOGGER_RMQ_URL"))
	sendThisErrorOnPriority(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	sendThisErrorOnPriority(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("LOGGER_RMQ_TOPIC"), // name
		false, // should the message be persistent? also queue will survive if the cluster gets reset
		false, // autodelete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
		false, // exclusive means I should get an error if any other consumer subsribes to this queue
		false, // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
		nil,   // arguments for more advanced configuration
	)
	sendThisErrorOnPriority(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	sendThisErrorOnPriority(err, "Failed to publish a message")
}
