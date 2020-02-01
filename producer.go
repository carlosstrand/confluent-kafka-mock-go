package kafkamock

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"testing"
)

type Producer interface {
	Close()
	Events() chan kafka.Event
	Flush(timeoutMs int) int
	GetFatalError() error
	GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error)
	Len() int
	OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error)
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
	ProduceChannel() chan *kafka.Message
	Purge(flags int) error
	QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error)
	SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error
	SetOAuthBearerTokenFailure(errstr string) error
	String() string
	TestFatalError(code kafka.ErrorCode, str string) kafka.ErrorCode
}

type ProducerMock struct {
	ProducedMessages []*ProducerMessage
	ProducedCount int
}

func (p * ProducerMock) Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error {
	p.ProducedMessages = append(p.ProducedMessages, &ProducerMessage{
		Message: msg,
	})
	p.ProducedCount += 1
	return nil
}

func (p * ProducerMock) LastProduced() *ProducerMessage {
	if len(p.ProducedMessages) == 0 {
		return nil
	}
	return p.ProducedMessages[len(p.ProducedMessages) - 1]
}

func (p * ProducerMock) ProducedAt(idx int) *ProducerMessage {
	if idx >= len(p.ProducedMessages) {
		return nil
	}
	return p.ProducedMessages[idx]
}

func (p * ProducerMock) AssertProducedCount(t *testing.T, n int) {
	if p.ProducedCount != n {
		t.Errorf("ProducedCount should be equal %d but it's equal %d", n, p.ProducedCount)
	}
}