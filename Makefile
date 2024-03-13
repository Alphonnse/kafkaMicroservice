get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-kafka-protoModel

generate-kafka-protoModel:
	mkdir -p internal/model/modelKafkaProto
	protoc --proto_path protoModel/kafkaModel \
	--go_out=internal/model/modelKafkaProto --go_opt=paths=source_relative \
	--go-grpc_out=internal/model/modelKafkaProto --go-grpc_opt=paths=source_relative \
	protoModel/kafkaModel/kafka.proto
