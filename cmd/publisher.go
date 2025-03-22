package cmd

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Rename publisher to Publisher to make it exported
func Publisher() {
	// Create an MQTT client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("go_mqtt_publisher")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to connect to the Broker:", token.Error())
		return
	}
	fmt.Println("Connected to the MQTT broker!")

	// Publish a message
	message := "Hello, MQTT!"
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	fmt.Println("Message published:", message)

	// Disconnect from the broker after a delay
	time.Sleep(2 * time.Second)
	client.Disconnect(250)
	fmt.Println("Disconnected from the MQTT broker!")
}
