package kafka

import (
	"context"

	"github.com/Alphonnse/kafkaMicroservice/internal/model/modelKafkaProto"
)

type KafkaProducer interface {
	PutMessage(context.Context, modelProto.Termometer) error
}

type KafkaConsumer interface {
	// string is the topic name
	GetMessage(context.Context, string) (modelProto.Termometer, error) 
}
