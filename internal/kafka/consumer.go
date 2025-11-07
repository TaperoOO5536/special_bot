package kafka

import (
	"fmt"
	"strings"

	"github.com/TaperoOO5536/special_bot/internal/handler"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
	orderHandler *handler.OrderHandler
	stop bool
}

func NewConsumer(address []string, topics []string, orderHandler handler.OrderHandler) (*Consumer, error) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),
		"group.id":          "mygroup",
	}

	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}
	if err = c.SubscribeTopics(topics, nil); err != nil {
    return nil, err
  }
	return &Consumer{consumer: c, orderHandler: &orderHandler}, nil
}

func (c *Consumer) Start() {
	fmt.Println("start consumer")
	for {
		if c.stop {
			break
		}
		kafkaMsg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			fmt.Println(err)
		}
		if kafkaMsg == nil {
			continue
		}
		var eventType string
    for _, h := range kafkaMsg.Headers {
    	if h.Key == "event-type" {
        eventType = string(h.Value)
        break
    	}
		}
		if eventType == "" {
      fmt.Println("missing event-type header")
      continue
    }
		fmt.Println(eventType)
		switch eventType {
		case "order.create":
			c.orderHandler.CreateOrder(kafkaMsg.Value)
		default:
			fmt.Println("unknown topic")
			continue
		}
	}
}

func (c *Consumer) Stop() error {
	fmt.Println("stop consumer")
	c.stop = true
	return c.consumer.Close()
}