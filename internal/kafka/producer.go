package kafka

import (
	"context"

	modelProto "github.com/Alphonnse/kafkaMicroservice/internal/model/modelKafkaProto"
)


// сюда прилетакт запрос из сервиса
// пут семмаге обращается к апи кафка и кладет туда 

func (p KafkaProducer) PutMessage(ctx context.Context, prod *modelProto.Termometer) error {
	return nil
}
