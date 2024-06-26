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

	"github.com/hamburgertrain/boostpi/internal/configuration"
	"github.com/hamburgertrain/elmobd"
)

// Establish contact with an ELM327 OBD-II reader
func Initialize(config configuration.Configuration) (*elmobd.Device, error) {
	dev, err := elmobd.NewDevice(config.Elm327DeviceLocation, config.Elm327Debug)
	if err != nil {
		return nil, err
	}

	return dev, nil
}

// Get version from an ELM327 OBD-II reader
func GetVersion(dev *elmobd.Device) (string, error) {
	version, err := dev.GetVersion()
	if err != nil {
		return "", err
	}

	return version, nil
}

// Get engine rpm from an ELM327 OBD-II reader
func GetEngineRpm(dev *elmobd.Device) (string, error) {
	rpm, err := dev.RunOBDCommand(elmobd.NewEngineRPM())
	if err != nil {
		return "", err
	}

	return rpm.ValueAsLit(), nil
}

// Get intake manifold pressure from an ELM327 OBD-II reader
func GetIntakeManifoldPressure(dev *elmobd.Device) (string, error) {
	imp, err := dev.RunOBDCommand(elmobd.NewIntakeManifoldPressure())
	if err != nil {
		return "", err
	}

	return imp.ValueAsLit(), nil
}

// Get mass airflow rate from an ELM327 OBD-II reader
func GetMassAirflowRate(dev *elmobd.Device) (string, error) {
	mafr, err := dev.RunOBDCommand(elmobd.NewMafAirFlowRate())
	if err != nil {
		return "", err
	}

	return mafr.ValueAsLit(), nil
}

// Get turbo compressor inlet pressure from an ELM327 OBD-II reader
func GetTurboCompressorInletPressure(dev *elmobd.Device) (string, error) {
	turboPressure, err := dev.RunOBDCommand(elmobd.NewTurbochargerCompressorInletPressure())
	if err != nil {
		return "", err
	}

	return turboPressure.ValueAsLit(), nil
}

// Get absolute barometric pressure from an ELM327 OBD-II reader
func GetAbsoluteBarometricPressure(dev *elmobd.Device) (string, error) {
	barometricPressure, err := dev.RunOBDCommand(elmobd.NewAbsoluteBarometricPressure())
	if err != nil {
		return "", err
	}

	return barometricPressure.ValueAsLit(), nil
}

// Get intake air temperature from an ELM327 OBD-II reader
func GetIntakeAirTemp(dev *elmobd.Device) (string, error) {
	intakeTemp, err := dev.RunOBDCommand(elmobd.NewIntakeAirTemperature())
	if err != nil {
		return "", err
	}

	return intakeTemp.ValueAsLit(), nil
}

// Check which commands are supported on a connected vehicle
// This loops on command '01C01' when not connected to a vehicle
func CheckSupportedCommands(dev *elmobd.Device) {
	supported, err := dev.CheckSupportedCommands()
	if err != nil {
		fmt.Println("Failed to check supported commands: ", err)
		return
	}

	allCommands := elmobd.GetSensorCommands()
	carCommands := supported.FilterSupported(allCommands)

	fmt.Printf("%d of %d commands supported:\n", len(carCommands), len(allCommands))

	for _, cmd := range carCommands {
		fmt.Printf("- %s supported\n", cmd.Key())
	}
}
