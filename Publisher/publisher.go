package publisher

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func sendThisErrorOnPriority(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Publish(url string, queue string, data ...interface{}) {
	conn, err := amqp.Dial(url)
	sendThisErrorOnPriority(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	sendThisErrorOnPriority(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // should the message be persistent? also queue will survive if the cluster gets reset
		false, // autodelete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
		false, // exclusive means I should get an error if any other consumer subsribes to this queue
		false, // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
		nil,   // arguments for more advanced configuration
	)
	sendThisErrorOnPriority(err, "Failed to declare a queue")

	publishData, err := json.Marshal(data[0])

	if err != nil {
		fmt.Print(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(publishData),
		})
	sendThisErrorOnPriority(err, "Failed to publish a message")
}
