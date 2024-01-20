// (eventually) Pull from ELM327 device and write boost data to an i2c LCD
package main

import (
	"log"
	"strconv"
	"time"

	"github.com/d2r2/go-i2c"
	"github.com/hamburgertrain/boostpi/display"
	"github.com/hamburgertrain/boostpi/elm327"
)

func main() {
	log.Println("Initializing connection to i2c display...")
	i2cConnection := display.Initialize()
	defer i2cConnection.Close()
	log.Println("Connection initialized")

	display.Reset(i2cConnection)

	display.LcdDisplayString(i2cConnection, "----BoostPi-----", 1, 0)
	display.LcdDisplayString(i2cConnection, "Now Loading.....", 2, 0)

	log.Println("Initializing connection to ELM327 device...")
	dev := elm327.Initialize()
	log.Println("Connection initialized")

	elm327.GetVersion(dev)
	elm327.GetEngineRpm(dev)
	elm327.GetMassAirflowRate(dev)
	elm327.GetIntakeManifoldPressure(dev)
	// This loops on command '01C01' when not connected to a vehicle
	//elm327.CheckSupportedCommands(dev)

	display.Clear(i2cConnection)

	log.Println("Simulating boost...")
	SimulateBoost(i2cConnection, 7, 18)
	time.Sleep(1 * time.Second)
	SimulateBoost(i2cConnection, 15, 21)

	log.Println("Turning display off")
	display.TurnOff(i2cConnection)
}

// Simulate boost climbing very crudely
func SimulateBoost(connection *i2c.I2C, start int, end int) {
	for i := start; i < end; i++ {
		displayString := strconv.Itoa(i) + " psi"
		display.LcdDisplayString(connection, displayString, 1, 0)
		time.Sleep(1 * time.Second)
		display.Clear(connection)
	}
}
