package publisher

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	//Make a connection
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	//Ccreate a channel
	ch, _ := conn.Channel()
	defer ch.Close()

	//Declare a queue
	q, err := ch.QueueDeclare(
		"hello", // name of the queue
		false,   // should the message be persistent? also queue will survive if the cluster gets reset
		false,   // autodelete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
		false,   // exclusive means I should get an error if any other consumer subsribes to this queue
		false,   // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
		nil,     // arguments for more advanced configuration
	)

	//Publish a message
	body := "hello world"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("Message: %s", body)

}
