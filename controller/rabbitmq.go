package controller

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type BookingRequest struct {
	Username string `json:"username"`
	Date     string `json:"date"`
	Charging string `json:"charging"`
	Indoor   string `json:"indoor"`
	Outdoor  string `json:"outdoor"`
}

func checkErrWithMsg(err error, msg string) {
	if err != nil {
		Error.Printf("%s: %s", msg, err)
	}
}

func RabbitMQSend(body []byte) {
	conn, err := amqp.Dial("amqp://guest:guest@47.97.82.144:5672/")
	checkErrWithMsg(err, "RabbitMQ...  failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	checkErrWithMsg(err, "RabbitMQ...  failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"ParkingLot", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	checkErrWithMsg(err, "RabbitMQ...  failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	logInfo("RabbitMQ...  sent %s", body)
	checkErrWithMsg(err, "RabbitMQ...  failed to publish a message")
}

func RabbitMQReceive() {
	conn, err := amqp.Dial("amqp://guest:guest@47.97.82.144:5672/")
	checkErrWithMsg(err, "RabbitMQ...  failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	checkErrWithMsg(err, "RabbitMQ...  failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"ParkingLot", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	checkErrWithMsg(err, "RabbitMQ...  failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	checkErrWithMsg(err, "RabbitMQ...  failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			logInfo("RabbitMQ...  received a message: ", d.Body)
			req := BookingRequest{}
			err := json.Unmarshal(d.Body, &req)
			checkErrWithMsg(err, "Failed to unmarshal the []byte")
			err = makeBooking(req)
			checkErrWithMsg(err, "Failed to make a Booking")
		}
	}()

	logInfo("RabbitMQ is waiting for messages...")
	<-forever
}
