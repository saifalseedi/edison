	

package main

import (
	 //	"time"
"encoding/json"
"fmt"
"time"
"os"
"gobot.io/x/gobot"
"gobot.io/x/gobot/drivers/gpio"
"gobot.io/x/gobot/platforms/intel-iot/edison"
MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Message struct {
	Name string
	Value bool

}

//Create structure for MQTT credentials
type Cred struct {
	user string
	password string
	clientId string
	topic string
	broker string

}

//Main looop
func main() {
	//MQTT credentials from target device
	mqtt_credentials := Cred{
		user: "<your user ID>",
		password: "<your password>",
		clientId: "<your client id>",
		topic: "<your topic>",
		broker: "<broker>"
	}

	//Instantiate a buzzer at digital pin 3 for ouput and board
	board := edison.NewAdaptor()
	buzzer := gpio.NewBuzzerDriver(board, "5")

	opts := MQTT.NewClientOptions()

	//Create broker through port 1883
	opts.AddBroker(mqtt_credentials.broker)

	opts.SetClientID(mqtt_credentials.clientId)
	opts.SetUsername(mqtt_credentials.user)
	opts.SetPassword(mqtt_credentials.password)
	opts.SetCleanSession(false)

	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
		})

	//Initialize the MQTT client
	client := MQTT.NewClient(opts)

	//Check for connection
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}


	if token := client.Subscribe(mqtt_credentials.topic + "cmd", 0 , nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//Work when robot is started
	work := func() {
		for {
			//Check for new messages from topic
			incoming := <-choke
			//print message and topic
			fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
			var m Message
			//convert from json to message structure
			json.Unmarshal([]byte(incoming[1]), &m)
			//print message value, ie true or false
			fmt.Println(m.Value)

			//if message value is true, turn buzzer on
			if m.Value == true {
				//start the buzzer
				buzzer.Tone(gpio.C4, gpio.Quarter)
				time.Sleep(10 * time.Millisecond)

				fmt.Printf("BUZZ")
		    //if message value is false, turn buzzer off
			} else if m.Value == false{
				fmt.Printf("NO BUZZ")
			}
		}
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{board},
		[]gobot.Device{buzzer},
		work,
		)

	robot.Start()
	client.Disconnect(250)
	fmt.Println("Sample Subscriber Disconnected")

}
