// Facilitates connection to and execution of commands on an ELM327 OBD-II reader
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
package elm327

import (
	"fmt"
	"log"

	"github.com/rzetterberg/elmobd"
)

var elm327DeviceLocation string = "/dev/ttyUSB0"
var elm327Debug bool = true

// Establish contact with an ELM327 OBD-II reader
func Initialize() *elmobd.Device {
	dev, err := elmobd.NewDevice(elm327DeviceLocation, elm327Debug)
	if err != nil {
		log.Fatal("Could not initialize ELM327 device:", err)
	}

	return dev
}

// Get version from an ELM327 OBD-II reader
func GetVersion(dev *elmobd.Device) {
	version, err := dev.GetVersion()
	if err != nil {
		log.Println("Failed to get version:", err)
		return
	}

	log.Println("Device has version:", version)
}

// Get engine rpm from an ELM327 OBD-II reader
func GetEngineRpm(dev *elmobd.Device) {
	rpm, err := dev.RunOBDCommand(elmobd.NewEngineRPM())
	if err != nil {
		log.Println("Failed to get RPM:", err)
		return
	}

	log.Printf("Engine spinning at %s RPMs\n", rpm.ValueAsLit())
}

// Get intake manifold pressure from an ELM327 OBD-II reader
func GetIntakeManifoldPressure(dev *elmobd.Device) {
	imp, err := dev.RunOBDCommand(elmobd.NewIntakeManifoldPressure())
	if err != nil {
		log.Println("Failed to get intake manifold pressure:", err)
		return
	}

	log.Printf("Intake manifold pressure is %s\n", imp.ValueAsLit())
}

// Get mass airflow rate from an ELM327 OBD-II reader
func GetMassAirflowRate(dev *elmobd.Device) {
	mafr, err := dev.RunOBDCommand(elmobd.NewMafAirFlowRate())
	if err != nil {
		log.Println("Failed to get mass airflow rate:", err)
		return
	}

	log.Printf("Mass airflow rate is %s\n", mafr.ValueAsLit())
}

// Check which commands are supported on a connected vehicle
func CheckSupportedCommands(dev *elmobd.Device) {
	supported, err := dev.CheckSupportedCommands()
	if err != nil {
		fmt.Println("Failed to check supported commands", err)
		return
	}

	allCommands := elmobd.GetSensorCommands()
	carCommands := supported.FilterSupported(allCommands)

	fmt.Printf("%d of %d commands supported:\n", len(carCommands), len(allCommands))

	for _, cmd := range carCommands {
		fmt.Printf("- %s supported\n", cmd.Key())
	}
}
