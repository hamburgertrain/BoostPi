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
	i2cConnection := display.Initialize()
	defer i2cConnection.Close()
	log.Println("Connection initialized")

	// Let's make sure we have blank slate
	display.Reset(i2cConnection)

	// Show loading text while contacting ELM327 device
	utilities.ShowLoadingText(i2cConnection)

	log.Println("Initializing connection to ELM327 device...")
	dev, err := elm327.Initialize()
	if err != nil {
		display.ShutdownDisplay(i2cConnection)
		log.Fatal("Could not initialize ELM327 device:", err)
	}
	log.Println("Connection initialized")

	// Initial probe of device
	version, err := elm327.GetVersion(dev)
	if err != nil {
		display.ShutdownDisplay(i2cConnection)
		log.Fatal("Error getting version:", err)
	}
	log.Println("Device has version:", version)

	// Clear our display of loading text before showing boost
	display.Clear(i2cConnection)

	// This loops on command '01C01' when not connected to a vehicle
	//elm327.CheckSupportedCommands(dev)

	log.Println("Simulating boost...")
	utilities.SimulateBoost(i2cConnection, 30)

	// Show that we can fetch information and display it
	// for true {
	// 	massAirflowRate, err := elm327.GetMassAirflowRate(dev)
	// 	if err != nil {
	// 		log.Printf("Failed to get mass airflow rate:", err)
	// 	}
	// 	log.Printf("Mass airflow rate is %s\n", massAirflowRate)

	// 	intakeManifoldPressure, err := elm327.GetIntakeManifoldPressure(dev)
	// 	if err != nil {
	// 		log.Printf("Failed to get intake manifold pressure:", err)
	// 	}
	// 	log.Printf("Intake manifold pressure is %s\n", intakeManifoldPressure)

	// 	display.LcdDisplayString(i2cConnection, massAirflowRate, 1, 0)
	// 	display.LcdDisplayString(i2cConnection, intakeManifoldPressure, 2, 0)

	// 	time.Sleep(1 * time.Second)
	// }

	log.Println("Turning display off")
	display.ShutdownDisplay(i2cConnection)
}
