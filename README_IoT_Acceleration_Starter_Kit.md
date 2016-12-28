# Intel® Edison and the IoT Acceleration Starter Kit

## Introduction

The [Intel® Edison](http://www.intel.com/content/www/us/en/do-it-yourself/edison.html) is a single-board computer with a Silvermont dual-core Intel Atom CPU, integrating WiFi, Bluetooth 4.0, 1 GB DDR, 4 GB eMMC flash memory and USB controllers. The 40 multiplexed GPIO pins, high computing power and connectivity capabilities allow rapid industrial IoT and fog computing prototyping.

In this repository, you will find tutorials and code examples of how to make use of the Edison as a sensor node with the [IoT Acceleration Starter Kit](http://www.iot-starterkit.de/), using two languages: Python and Arduino.

![](./assets/edison_intro_pic_dell.jpg)

This document provides resources on how to get started with the Edison, including the mechanical assembly. **If your board has already been assembled,
then please navigate to one of the available tutorials to continue:**

-  [Python](https://github.com/relayr/edison/tree/master/python)
-  [Arduino](https://github.com/relayr/edison/tree/master/arduino)

## Requirements

The following hardware is required:

-  [Dell Edge Gateway 5100](http://www.dell.com/us/business/p/dell-edge-gateway-5100/pd)
-  [Intel® Edison Arduino Breakout Kit](https://www.arduino.cc/en/ArduinoCertified/IntelEdison#toc3)
- Micro-USB data cable

A computer is required to connect the Edison board for flashing and configuration.

**NOTE:** The additional hardware required for the code examples is specified on the Arduino and Python tutorials. All necessary parts are included in the [IoT Acceleration Starter Kit](http://www.iot-starterkit.de/).

## Installation & Configuration

### Setting Up the Hardware

Assemble the Arduino Expansion Board found in your Dell Starter Kit according to the directions in the [Intel® Edison guide](https://software.intel.com/en-us/node/628221).

Before continuing, make sure that your board looks like this:

![Edison assembled board](./assets/edison_assembled_board.jpg)

After that, plug the Grove Base Shield onto your Arduino Expansion Board as shown below, and you'll be ready to go:

![Edison with the base shield](./assets/edison_base_shield.jpg)

### Setting Up the Software

Once our board is ready, it's time to configure the software of the Edison according to the programming language of your choice. **Before proceeding with
the Vertex integration,** select one of the following tutorials, and complete all the steps:

-  [Python](https://github.com/relayr/edison/tree/master/python)
-  [Arduino](https://github.com/relayr/edison/tree/master/arduino)

## Association with Vertex

In order to associate your Edison with Vertex, you must first set up the Edison software according to the programming language of your choice, as explained in the previous step.

Once you've done that, then see the [Vertex guide](http://docs.relayr.io/iot-starter-kits/dsk/vertex#Adding-devices-to-your-vertex) for instructions on how to perform the device association.

## License

Copyright (C) 2016 relayr GmbH, Klemen Lilija <klemen@relayr.io>, Brian Lemke <brian@relayr.io>, Antonio Almeida <antonio@relayr.io>, Jaime González-Arintero <jaime@relayr.io>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

Except as contained in this notice, the name(s) of the above copyright holders shall not be used in advertising or otherwise to promote the sale, use or
other dealings in this Software without prior written authorization.

THE SOFTWARE IS PROVIDED "AS IS," WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.