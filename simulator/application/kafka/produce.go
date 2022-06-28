package kafka

import (
	"encoding/json"
	"github/limaleandro1999/simulator/application/route"
	"github/limaleandro1999/simulator/infra/kafka"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJSONPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KAFKA_PRODUCE_TOPIC"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
