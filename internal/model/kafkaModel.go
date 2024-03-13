package model

import "time"

type KafkaPoduserGlobalModel struct {
	Temperature float32
	Humidity    float32
	Location    string
	Timestamp   *time.Time
}
