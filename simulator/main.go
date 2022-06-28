package main

import (
	"fmt"
	kafka2 "github/limaleandro1999/simulator/application/kafka"
	"github/limaleandro1999/simulator/infra/kafka"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		go kafka2.Produce(msg)
		fmt.Println(string(msg.Value))
	}
	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringJson, _ := route.ExportJSONPositions()
	// fmt.Println(stringJson[0])
}
