package main

import (
	"github.com/streadway/amqp"
	"log"
)

// We also need an helper function to check the return value for each amqp call
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接到RabbitMQ Server
	conn, err := amqp.Dial("amqp://user_mmr:123456@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 创建channel，channel为驻留程序
	ch, err := conn.Channel()
	failOnError(err, "Failed to Open a channel")
	defer ch.Close()
	// 我们需要声明发送队列，那么我们可以发送消息到这个队列上
	// 队列的创建时幂等的，只有当不存在，才会创建队列，且队列名是唯一的
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	body := "Hey,judy!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
