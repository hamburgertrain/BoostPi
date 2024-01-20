// (eventually) Pull from ELM327 device and write boost data to an i2c LCD
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/d2r2/go-i2c"
	"github.com/hamburgertrain/boostpi/display"
	"github.com/hamburgertrain/boostpi/elm327"
)

func main() {
	fmt.Println("Initializing connection to ELM327 device...")
	elm327.ContactElm327Device()
	fmt.Println("Connection initialized")

	fmt.Println("Initializing connection to i2c display...")
	// Create new connection to I2C bus on 2 line with address 0x27
	i2cConnection := display.Initialize()
	// Free I2C connection on exit
	defer i2cConnection.Close()
	fmt.Println("Connection initialized")

	fmt.Println("Resetting display...")
	display.Reset(i2cConnection)
	time.Sleep(1 * time.Second)
	fmt.Println("Display reset")

	fmt.Println("Writing to display...")
	display.LcdDisplayString(i2cConnection, "Hello world!", 1, 0)
	fmt.Println("Writing done.")

	time.Sleep(5 * time.Second)

	fmt.Println("Clearing display...")
	display.Clear(i2cConnection)
	fmt.Println("Clear")

	fmt.Println("Simulating boost...")
	SimulateBoost(i2cConnection)

	fmt.Println("Turning display off")
	display.TurnOff(i2cConnection)
}

// Simulate boost climbing very crudely
func SimulateBoost(connection *i2c.I2C) {
	for i := 0; i < 30; i++ {
		display.LcdDisplayString(connection, strconv.Itoa(i), 1, 0)
		time.Sleep(1 * time.Second)
		display.Clear(connection)
		time.Sleep(1 * time.Second)
	}
}
