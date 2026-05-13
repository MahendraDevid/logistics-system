package kafka

import "log"

type Producer struct {
}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) PublishDeliveredEvent(awb string) {

	log.Printf(
		"event PackageDelivered published for AWB %s",
		awb,
	)
}