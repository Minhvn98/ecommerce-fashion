package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/Minhvn98/ecommerce-fashion/services/rabbitmq/common"
	"github.com/streadway/amqp"
)

var Rmq common.RabbitMQ
var Chanel *amqp.Channel

func RunRabbitMQ() {
	var wg sync.WaitGroup
	// Rabbitmq
	Rmq = common.RabbitMQ{
		ConnectionString: "amqp://guest:guest@localhost:5672/#/",
	}
	Rmq.CreateConnection()
	defer Rmq.Close()
	Chanel := Rmq.GetChannel()
	q, err := Chanel.QueueDeclare(
		"UploadProduct", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	msgs := common.Consume(Chanel, q.Name)
	common.FailOnError(err, "Failed to register a consumer")
	wg.Add(1)
	go func() {
		for d := range msgs {
			myString := string(d.Body[:])
			fmt.Println(myString)
		}
	}()
	wg.Wait()

	common.Publish(Chanel, "UploadProduct", "uk")

}
