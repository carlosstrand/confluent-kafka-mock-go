package main

import (
	"encoding/json"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"testing"
)

type CustomEvent struct {
	Foo string `json:"foo"`
}

type MyApp struct {
	producer Producer
}

func (a * MyApp) sendMessage() {
	event := CustomEvent{
		Foo: "My Event!",
	}
	value, _ := json.Marshal(event)
	topic := "ms.myapp.messages"
	_ = a.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			Partition: 2,
			Offset: 4,
		},
		Value:          value,
	}, nil)
}

func TestProduce(t *testing.T) {
	p := ProducerMock{}
	app := MyApp{&p}
	p.AssertProducedCount(t, 0)
	app.sendMessage()
	p.AssertProducedCount(t, 1)
	p.LastProduced().AssertTopic(t, "ms.myapp.messages")
	p.LastProduced().AssertPartition(t, 2)
	p.LastProduced().AssertOffset(t, 4)
	p.LastProduced().AssertValueEquals(t, CustomEvent{
		Foo: "My Event!",
	})
}