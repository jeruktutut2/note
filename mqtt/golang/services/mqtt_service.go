package services

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttService interface {
	SendMessage(message string) (response string)
}

type MqttServiceImplementation struct {
	Client mqtt.Client
}

func NewMqttService(client mqtt.Client) MqttService {
	return &MqttServiceImplementation{
		Client: client,
	}
}

func (service *MqttServiceImplementation) SendMessage(message string) (response string) {
	token := service.Client.Publish("test/topic", 0, false, message)
	token.Wait()
	fmt.Println("sending message:", message)
	response = "success"
	return
}
