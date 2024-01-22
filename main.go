// Pull boost data from an ELM327 device and write it to an i2c LCD
// Copyright (C) 2024 Tyler Bialoblocki
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"log"

	"github.com/hamburgertrain/boostpi/display"
	"github.com/hamburgertrain/boostpi/elm327"
	"github.com/hamburgertrain/boostpi/utilities"
)

// Application entrypoint
func main() {
	log.Println("Initializing connection to i2c display...")
	i2cConnection, err := display.Initialize()
	if err != nil {
		log.Fatal("Could not initialize i2c device:", err)
	}
	defer i2cConnection.Close()
	log.Println("Connection initialized")

	// Let's make sure we have blank slate
	display.Reset(i2cConnection)

	// Show loading text while contacting ELM327 device
	display.LcdDisplayString(i2cConnection, "----BoostPi-----", 1, 0)
	display.LcdDisplayString(i2cConnection, "----Loading-----", 2, 0)

	log.Println("Initializing connection to ELM327 device...")
	dev, err := elm327.Initialize()
	if err != nil {
		display.TurnBacklightOff(i2cConnection)
		display.TurnDisplayOff(i2cConnection)
		log.Fatal("Could not initialize ELM327 device:", err)
	}
	log.Println("Connection initialized")

	// Clear our display of loading text before showing boost
	display.Clear(i2cConnection)

	// This loops on command '01C01' when not connected to a vehicle
	//elm327.CheckSupportedCommands(dev)

	// We want to loop over these and display them
	massAirflowRate, err := elm327.GetMassAirflowRate(dev)
	if err != nil {
		log.Printf("Failed to get mass airflow rate:", err)
	}
	log.Printf("Mass airflow rate is %s\n", massAirflowRate)

	intakeManifoldPressure, err := elm327.GetIntakeManifoldPressure(dev)
	if err != nil {
		log.Printf("Failed to get intake manifold pressure:", err)
	}
	log.Printf("Intake manifold pressure is %s\n", intakeManifoldPressure)

	log.Println("Simulating boost...")
	utilities.SimulateBoost(i2cConnection, 30)

	// Show that we can fetch information and display it
	// for true {
	// 	version, err := elm327.GetVersion(dev)
	// 	if err != nil {
	// 		log.Println("Error getting version:", err)
	// 	}
	// 	log.Println("Device has version:", version)
	// 	display.LcdDisplayString(i2cConnection, version, 1, 0)
	// 	time.Sleep(1 * time.Second)
	// }

	log.Println("Turning display off")
	display.TurnBacklightOff(i2cConnection)
	display.TurnDisplayOff(i2cConnection)
}
