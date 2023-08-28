package nsq

import (
	"fmt"
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

var (
	NsqProducer *nsq.Producer
)

func StartProducer() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(os.Getenv("NSQ_ADDRESS"), config)
	if err != nil {
		log.Fatal(err)
	}

	NsqProducer = producer
	fmt.Println("NSQ producer started!")
}
