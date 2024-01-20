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
	"math/rand"
	"strconv"
	"time"

	"github.com/d2r2/go-i2c"
	"github.com/hamburgertrain/boostpi/display"
	"github.com/hamburgertrain/boostpi/elm327"
)

// Application entrypoint
func main() {
	log.Println("Initializing connection to i2c display...")
	i2cConnection := display.Initialize()
	defer i2cConnection.Close()
	log.Println("Connection initialized")

	display.Reset(i2cConnection)

	display.LcdDisplayString(i2cConnection, "----BoostPi-----", 1, 0)
	display.LcdDisplayString(i2cConnection, "----Loading-----", 2, 0)

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
	SimulateBoost(i2cConnection, 30)

	log.Println("Turning display off")
	display.TurnOff(i2cConnection)
}

// Simulate boost climbing very crudely
func SimulateBoost(connection *i2c.I2C, end int) {
	for i := 0; i < end; i++ {
		rn := (rand.Float64() * 5) + 5
		s := strconv.FormatFloat(rn, 'f', 2, 64)
		displayString := s + " psi"

		display.LcdDisplayString(connection, displayString, 1, 0)
		time.Sleep(1 * time.Second)

		display.Clear(connection)
	}
}
