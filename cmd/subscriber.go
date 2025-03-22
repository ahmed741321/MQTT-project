// Description: A simple MQTT subscriber that subscribes to a topic and prints the received messages to the console.

package cmd

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscriber() {
	// Create an MQTT client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("go_mqtt_subscriber")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to connect to the Broker:", token.Error())
		os.Exit(1)
	}

	// Subscribe to the topic
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("📩 Message received: %s\n", msg.Payload())
	})

	// Keep the subscriber running
	select {}
}
