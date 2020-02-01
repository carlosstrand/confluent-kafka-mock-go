# confluent-kafka-mock-go

Kafka Mocks that makes it easy to test applications that use confluent-kafka-go


### Quickstart

In your application, receive our Producer interface (or create one) as an argument:

```go
type MyApp struct {
  producer kafkamock.Producer
}
```

In a real-world case, you just need to create a producer and use it in your app:

``` go
p := kafka.NewProducer(libkafka.ConfigMap{
  "bootstrap.servers": "localhost",
  "group.id": "mygroup",
  "broker.address.family": "v4",
})

app := MyApp{&p}
  
// Sending some message
topic := "ms.myapp.messages"
app.producer.Produce(&kafka.Message{
  TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
  Value: "Some value here!",
}, nil)
```

In a test case, you just need to create a producer mock:

```go
func TestProduce(t *testing.T) {
  p := ProducerMock{}
  app := MyApp{&p}
  p.AssertProducedCount(t, 0)
  app.sendMessage()
  p.AssertProducedCount(t, 1)
  p.LastProduced().AssertTopic(t, "ms.myapp.messages")
  p.LastProduced().AssertPartition(t, 2)
  p.LastProduced().AssertOffset(t, 4)
  p.LastProduced().AssertValueEquals(t, "Some value here!")
}
```


### Contributing to this package
We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer


Feel free to open an Issue or a PR! 🚀
```go
