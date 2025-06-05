package mq

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	mqtp "pms.pkg/transport/mq"
	notifiermq "pms.pkg/transport/mq/notifier"
	"pms.pkg/utils"
)

var (
	testRouting mqtp.QueueRoute = "test"
)

const (
	testDSN      = "amqp://guest:guest@0.0.0.0:5672/"
	testExchange = "direct-exchange"
	testQueue    = "test-queue"
)

type TestMessage struct {
	Content string
}

func (m TestMessage) RoutingKey() mqtp.QueueRoute {
	return testRouting
}

func TestRabbitMQIntegration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := Config{
		DSN:      testDSN,
		Exchange: testExchange,
	}

	log, _ := zap.NewDevelopment()
	sugar := log.Sugar()

	subscriber, err := NewSubscriber(ctx, SubscriberOpts{
		Queue:  testQueue,
		Routes: []mqtp.QueueRoute{testRouting},
		Config: conf,
		Log:    sugar,
	})
	assert.NoError(t, err, "failed to create subscriber")

	msgs, err := subscriber.Consume(ctx)
	assert.NoError(t, err, "failed to start consumer")

	received := make(chan mqtp.Queueable, 1)

	go func() {
		for msg := range msgs {
			var receivedMsg TestMessage
			err := json.Unmarshal(msg.Body, &receivedMsg)
			if err != nil {
				t.Errorf("failed to decode message: %v", err)
			} else {
				t.Log("received msg", receivedMsg)
				received <- receivedMsg
			}
			msg.Ack(false)
		}
	}()

	publisher, err := NewPublisher(ctx, PublisherOpts{
		Queue:  "test-queue",
		Config: conf,
		Logger: sugar,
	})
	t.Log("publisher state: ", publisher.ConnState())
	assert.NoError(t, err, "failed to create publisher")

	testMsg := TestMessage{Content: "Hello, RabbitMQ!"}
	err = publisher.Publish(ctx, testMsg)
	assert.NoError(t, err, "failed to publish message")

	select {
	case msg := <-received:
		t.Log(utils.JSON(msg))
	case <-time.After(5 * time.Second):
		t.Fatal("timeout: no message received")
	}

	publisher.Close()
	subscriber.Close()
}

func Test_ProduceGreetMessage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := Config{
		DSN:      testDSN,
		Exchange: testExchange,
	}

	publisher, err := NewPublisher(ctx, PublisherOpts{
		Queue:  "notifier",
		Config: conf,
		Logger: zap.NewNop().Sugar(),
	})
	if err != nil {
		t.Fatal(err)
	}
	greetMsg := notifiermq.GreetMessage{
		Name: "Danil",
		MetaData: notifiermq.MetaData{
			ToEmail: "danil.li24x@gmail.com",
		},
	}
	if err := publisher.Publish(ctx, greetMsg); err != nil {
		t.Fatal(err)
	}
}
