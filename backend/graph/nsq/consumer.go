package nsq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type messageHandler struct{}

type Message struct {
	Name      string
	Content   interface{}
	Timestamp time.Time
}

var (
	ConsumerMessage Message
)

func StartConsumer(topic string, channel string) {
	config := nsq.NewConfig()

	config.MaxAttempts = 10
	config.MaxInFlight = 5

	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&messageHandler{})
	consumer.ConnectToNSQLookupd(os.Getenv("NSQ_LOOKUPD"))

	fmt.Println("NSQ consumer started!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumer.Stop()
	fmt.Println(ConsumerMessage)
}

func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	var message Message

	if err := json.Unmarshal(m.Body, &message); err != nil {
		log.Println("error when Unmarshal() the message body, err:", err)
		return err
	}
	ConsumerMessage = message
	log.Println(message)

	return nil
}
