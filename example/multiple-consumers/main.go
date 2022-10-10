package main

import (
	"github.com/k0kubun/pp"
	"kafka-cronsumer"
	"kafka-cronsumer/internal/config"
	"kafka-cronsumer/model"
)

func main() {
	first := getConfig("config-1")
	var firstConsumerFn kafka_cronsumer.ConsumeFn = func(message model.Message) error {
		pp.Printf("First Consumer > Message received: %s\n", string(message.Value))
		return nil
	}
	firstHandler := kafka_cronsumer.NewKafkaHandler(first.Kafka, firstConsumerFn, true)
	firstHandler.Start(first.Kafka.Consumer)

	second := getConfig("config-2")
	var secondConsumerFn kafka_cronsumer.ConsumeFn = func(message model.Message) error {
		pp.Printf("Second Consumer > Message received: %s\n", string(message.Value))
		return nil
	}
	secondHandler := kafka_cronsumer.NewKafkaHandler(second.Kafka, secondConsumerFn, true)
	secondHandler.Start(first.Kafka.Consumer)

	select {} // block main goroutine (we did to show it by on purpose)
}

func getConfig(configName string) *config.ApplicationConfig {
	cfg, err := config.New("./example/multiple-consumers", configName)
	if err != nil {
		panic("application config read failed: " + err.Error())
	}
	cfg.Print()
	return cfg
}