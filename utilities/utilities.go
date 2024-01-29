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
	"strconv"
	"time"

	"github.com/d2r2/go-i2c"
	"github.com/hamburgertrain/boostpi/display"
	"github.com/hamburgertrain/boostpi/elm327"
	"github.com/hamburgertrain/elmobd"
)

const (
	psiConversion float64 = 0.145038
)

// Loop over values and display them
func GetAndDisplayValues(connection *i2c.I2C, obdDevice *elmobd.Device) {
	for {
		barometricPressure, err := elm327.GetAbsoluteBarometricPressure(obdDevice)
		if err != nil {
			display.ShowErrorAndShutdown(connection)
			log.Fatal("Failed to get barometric pressure:", err)
		}

		intakeManifoldPressure, err := elm327.GetIntakeManifoldPressure(obdDevice)
		if err != nil {
			display.ShowErrorAndShutdown(connection)
			log.Fatal("Failed to get intake manifold pressure:", err)
		}

		parsedManifoldPressure, err := strconv.ParseUint(intakeManifoldPressure, 10, 32)
		if err != nil {
			display.ShowErrorAndShutdown(connection)
			log.Fatal("Failed to convert intake manifold pressure:", err)
		}

		parsedBarometricPressure, err := strconv.ParseUint(barometricPressure, 10, 32)
		if err != nil {
			display.ShowErrorAndShutdown(connection)
			log.Fatal("Failed to convert barometric pressure:", err)
		}

		var calculatedManifoldPressure uint64 = 0
		if parsedManifoldPressure > parsedBarometricPressure {
			calculatedManifoldPressure = (parsedManifoldPressure - parsedBarometricPressure)
		}

		// Do our boost calculation and convert to psi
		calculatedBoostPressure := (float64(calculatedManifoldPressure) * psiConversion)

		// We don't want to display negative boost pressure
		if calculatedBoostPressure < 0 {
			calculatedBoostPressure = 0
		}

		stringFloat := strconv.FormatFloat(calculatedBoostPressure, 'f', 2, 64)

		intakePressureDisplay := stringFloat + " psi"

		display.LcdDisplayString(connection, intakePressureDisplay, 1, 0)
		display.LcdDisplayString(connection, intakeManifoldPressure+" : "+barometricPressure, 2, 0) // Display for debug, this is in kPa

		time.Sleep(1 * time.Second)
	}
}
