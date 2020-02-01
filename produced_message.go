package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ProducerMessage struct {
	Message *kafka.Message
}

func (m * ProducerMessage) AssertValueEquals(t *testing.T, v interface{}) {
	value, _ := json.Marshal(v)
	isEqual := bytes.Compare(m.Message.Value, value) == 0
	if !isEqual {
		t.Errorf("Produced message (%s) is not equal as expected (%s)", m.Message.Value, value)
	}
}

func (m * ProducerMessage) AssertTopic(t *testing.T, topic string) {
	tc := m.Message.TopicPartition.Topic
	if *tc != topic {
		t.Errorf("Produced topic (%s) is not equal as expected (%s)", *tc, topic)
	}
}

func (m * ProducerMessage) AssertPartition(t *testing.T, partition int32) {
	p := m.Message.TopicPartition.Partition
	if p != partition {
		t.Errorf("Produced partition (%d) is not equal as expected (%d)", p, partition)
	}
}

func (m * ProducerMessage) AssertOffset(t *testing.T, offset kafka.Offset) {
	o := m.Message.TopicPartition.Offset
	if o != offset {
		t.Errorf("Produced offset (%d) is not equal as expected (%d)", o, offset)
	}
}
