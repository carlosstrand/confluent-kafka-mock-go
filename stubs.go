package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (p * ProducerMock) Close() {
	return
}

func (p * ProducerMock) Events() chan kafka.Event {
	return nil
}

func (p *ProducerMock) Flush(timeoutMs int) int {
	return 0
}

func (p *ProducerMock) GetFatalError() error {
	return nil
}

func (p *ProducerMock) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error) {
	return nil, nil
}

func (p *ProducerMock) Len() int {
	return 0
}

func (p *ProducerMock) OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error) {
	return nil, nil
}

func (p *ProducerMock) ProduceChannel() chan *kafka.Message {
	return nil
}

func (p *ProducerMock) Purge(flags int) error {
	return nil
}

func (p * ProducerMock) QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error) {
	return 0, 0, nil
}

func (p * ProducerMock) SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error {
	return nil
}

func (p * ProducerMock) SetOAuthBearerTokenFailure(errstr string) error {
	return nil
}

func (p * ProducerMock) String() string {
	return ""
}

func (p * ProducerMock) TestFatalError(code kafka.ErrorCode, str string) kafka.ErrorCode {
	return 0
}
