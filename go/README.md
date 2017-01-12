# Intel Edison and the IoT Acceleration Starter Kit (Go Edition)

## Introduction

This document details all the necessary steps required
to make use of the [Intel Edison](http://www.intel.com/content/www/us/en/do-it-yourself/edison.html)
as a sensor node with the [IoT Acceleration Starter Kit](http://www.iot-starterkit.de/).

This document refers **only** to using **Go** on the Edison. See [here](https://github.com/relayr/edison/tree/master/Arduino) for using Arduino instead. See [here](https://github.com/relayr/edison/tree/master/Python) for using Python instead

## Requirements

The following hardware is required:

 * [Intel Edison Arduino breakout kit](https://www.iot-starterkit.de/fileadmin/media/pdf/iot/edisonarduino_hg_331191007_2.pdf)
 * [Seeedstudio Grove Base Shield](http://wiki.seeedstudio.com/wiki/Grove_-_Base_shield_v2)
 * Seeedstudio Sensor Kit

A computer is required to connect to the Edison board for flashing and configuring.

## Installation & configuration

### Setting up the hardware

Assemble the Arduino Expansion Board found in your IoT Acceleration Starter Kit according to the directions in the [Intel Edison guide](https://software.intel.com/en-us/node/628221).

After that, plug the Grove Base Shield onto your Arduino Expansion Board as shown below, and you'll be ready to go:

![Edison with the base shield](../assets/edison_base_shield.jpg)

### Setting up the firmware

To flash the latest firmware on the Edison board, use the *Intel® Edison Board Configuration Tool*
([Mac OS X](https://software.intel.com/en-us/get-started-edison-osx-step2),
[Linux](https://software.intel.com/en-us/get-started-edison-linux-step2),
[Windows](https://software.intel.com/en-us/get-started-edison-windows-step2)).
The official firmware is an operating system provided by the [Yocto Project](https://www.yoctoproject.org/), which is a Linux flavor geared towards embedded systems.

The setup wizard will guide you through the firmware flashing process, SSH configuration and WiFi configuration. Start with flashing the firmware and continue with SSH and WiFi configuration. When a step is completed, a green checkmark is shown next to the setup options in the wizard.

![IP Address](./assets/board-configuration-tool.png)


### Connecting your personal computer with the Intel Edison

When done with firmware flashing and the basic configurations of the Edison board, we can SSH into Intel Edison board and do further configurations by using its Linux shell. The Edison's IP address is found in the WiFi section of the **Intel Edison Board Configuration Tool**.

Alternatively, we can find the IP address of the board **without** the Intel configuration wizard by using any of the following methods:

 1. Log in to your router/access point and find the IP address assigned to the Edison board.
 2. Set up [mDNS](http://www.multicastdns.org/) on your Intel Edison.
 3. Set up the board using a **static** IP.
 4. Find the IP address of the board using a scanner such as [`nmap`](https://nmap.org/).

Once we know the IP of the Intel Edison board we can execute the following command which will log us in the Edison board using SSH:

```shell
ssh root@<edison's-IP-address>
```

When prompted for a password use the one that was set in the security settings of the Intel Edison Board Configuration Tool. If everything worked properly, we should now be
logged in to the board as **root**.


###Install Go
To install Go, follow the [official installation instructions](https://golang.org/doc/install). 

Download the go1.7.4.linux-386.tar.gz version for the Intel Edison. Navigate to the file location and move it to the Intel Edison,

```shell
scp go1.7.4.linux-386.tar.gz root@<IP of your device>:/usr/local
```
extract the package,

```shell
mkdir usr/local
tar -C usr/local -xzf home/root/go1.7.4.linux-386.tar.gz
```

and finally add /usr/local/go/bin to the PATH environment variable via

```shell
export PATH=$PATH:/usr/local/go/bin
```

### Install git-perl tools
The lightweight version of Git that comes with the Intel Edison does not support the submodules used by Gobot. Installing git-perltools is necessary.

```shell
opkg install git-perltools
```

### Install Gobot
To use Go with the Intel Edison, use the Gobot framework for drivers and adaptors on the board. To install Gobot and its required dependencies, execute the following in your computer terminal:

``` shell
 go get -d -u gobot.io/x/gobot/...
```

### Clone the git repository

Once you are logged into the Intel Edison and are able to interact with its Linux shell, you can clone the git repository by executing the following command:

```shell
git clone https://github.com/relayr/edison
```

From then on you can find the folder with the Go code examples `~/edison/go/examples` of the Edison board. However, before running any of them we first have to install all the necessary packages.

##Code Examples
### Example 1 (blink.go)

The `blink.go` example is a *Hello world* script which will toggle the LED on the Edison board every second. 

```shell
go run blink.go
```

If the LED on the board started blinking, then you can move on to the next example, in which we will connect a motion sensor to the Edison board. 

### Example 2 (temperature.go)

The `temperature.go` example uses a Grove Temperature Sensor. The sensor outputs a analog value for the temperature in Celsius that is then sent to the relayr Cloud.

First, prepare the hardware by connecting the *Grove temperature sensor* to the **Analog Pin 3 (A3)**.

Next, modify the `temperature.go` script in the temperature folder with the MQTT credentials of the device you've created on the relayr Dashboard. 

```go
# MQTT credentials.
	mqtt_credentials := Cred{
    	user: "<your user ID>",
    	password: "<your password>",
    	clientId: "<your client id>",
    	topic: "<your topic>",
	}
```

```shell
go run temperature.go
```

If the MQTT client connects, you can observe the changes on the relayr Dashboard on the temperature sensor.

### Example 3 (buzzer.go)

The `buzzer.go` example shows you how to receive commands from the relayr Cloud. The command received by Intel Edison will remotely turn on or off a piezo buzzer.

First prepare the hardware by connecting the [grove
buzzer](http://wiki.seeedstudio.com/wiki/Grove_-_Buzzer) to the **Digital pin 5 (D5)**.

![D5 pin](../assets/d5-pin.jpg)


Next, modify the `buzzer.go` script in the example folder with the MQTT credentials of the device you've created on the relayr Dashboard. 

```go
# MQTT credentials.
	mqtt_credentials := Cred{
    	user: "<your user ID>",
    	password: "<your password>",
    	clientId: "<your client id>",
    	topic: "<your topic>",
	}
```

```shell
go run buzzer.go
```
If the MQTT client connects, you can observe the changes on the relayr Dashboard on the buzzer.

Now the Intel Edison is listening to the messages from the relayr cloud. You can control the buzzer by pressing **True** or **False** in the **buzzer**
widget on the relayr Dashboard. If you set it up correctly, you'll hear a buzzing sound.
