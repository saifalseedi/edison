
package main

import (
"fmt"
MQTT "github.com/eclipse/paho.mqtt.golang"
"time"
"strconv"
"gobot.io/x/gobot"
"gobot.io/x/gobot/drivers/aio"
"gobot.io/x/gobot/platforms/intel-iot/edison"
)

//Create structure for MQTT credentials
type Cred struct {
	user string
	password string
	clientId string
	topic string
	broker string
}

//publish the payload as a json message to the 'data' MQTT topic
func pub_value(topic string, client MQTT.Client, value string, meaning_name string){
	var jsonprep string = `{"meaning":"`+meaning_name+`","value":"`+value+`"}`
	token := client.Publish(topic, byte(0), false, jsonprep)
	token.Wait()
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

	opts := MQTT.NewClientOptions()
	
	//Create broker through port 1883
	opts.AddBroker(mqtt_credentials.broker)
	opts.SetClientID(mqtt_credentials.clientId)
	opts.SetUsername(mqtt_credentials.user)
	opts.SetPassword(mqtt_credentials.password)
	opts.SetCleanSession(false)
	
	//Initialize the MQTT client
	client := MQTT.NewClient(opts)

	//Check for connection
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//Instantiate a temperature sensor object at analog pin A0 for input and board
	board := edison.NewAdaptor()
	sensor := aio.NewGroveTemperatureSensorDriver(board, "0")
	
	work := func() {
		gobot.Every(500*time.Millisecond, func() {
            		//Form the two values required in payload: meaning and value
			value := strconv.FormatFloat(sensor.Temperature(), 'E', -1, 64)
			meaning_name := "temperature"
                 	//Send payload values to publish to MQTT topic
			pub_value(mqtt_credentials.topic + "data", client, value, meaning_name)
			})
	}

    //Instatiate new gobot configuration
	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
		)

	robot.Start()

	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")	
}
