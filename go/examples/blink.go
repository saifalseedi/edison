package main

import (
"time"
"gobot.io/x/gobot"
"gobot.io/x/gobot/drivers/gpio"
"gobot.io/x/gobot/platforms/intel-iot/edison"
)

func main() {
        //Instatiate the edison adaptor and the led at GPIO 13
        e := edison.NewAdaptor()
        led := gpio.NewLedDriver(e, "13")
        
        work := func() { 
                gobot.Every(1*time.Second, func() {
                        //Turn the LED off and on every 0.5 seconds
                        led.Toggle()
                        })
        }

        robot := gobot.NewRobot("blinkBot",
                []gobot.Connection{e},
                []gobot.Device{led},
                work,
                )

        robot.Start()
}