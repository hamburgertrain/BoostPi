// Houses utility functions
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
package utilities

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/d2r2/go-i2c"
	"github.com/hamburgertrain/boostpi/display"
	"github.com/hamburgertrain/boostpi/elm327"
	"github.com/rzetterberg/elmobd"
)

// Loop over values and display them
func GetAndDisplayValues(connection *i2c.I2C, obdDevice *elmobd.Device) {
	for true {
		massAirflowRate, err := elm327.GetMassAirflowRate(obdDevice)
		if err != nil {
			display.ShowErrorAndShutdown(connection)
			log.Fatal("Failed to get mass airflow rate:", err)
		}
		log.Printf("Mass airflow rate is %s\n", massAirflowRate)

		intakeManifoldPressure, err := elm327.GetIntakeManifoldPressure(obdDevice)
		if err != nil {
			display.ShowErrorAndShutdown(connection)
			log.Fatal("Failed to get intake manifold pressure:", err)
		}
		log.Printf("Intake manifold pressure is %s\n", intakeManifoldPressure)

		display.LcdDisplayString(connection, massAirflowRate, 1, 0)
		display.LcdDisplayString(connection, intakeManifoldPressure, 2, 0)

		time.Sleep(1 * time.Second)
	}
}

// Simulate boost numbers very crudely
func SimulateBoost(connection *i2c.I2C, end int) {
	for i := 0; i < end; i++ {
		stringfloat := getRandomFloatAsString()
		displayString := stringfloat + " psi"

		display.LcdDisplayString(connection, displayString, 1, 0)
		time.Sleep(1 * time.Second)
	}
}

// Get a random float and convert it to a string for display
func getRandomFloatAsString() string {
	randomFloat := (rand.Float64() * 5) + 5
	stringfloat := strconv.FormatFloat(randomFloat, 'f', 2, 64)
	return stringfloat
}
