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
	"github.com/d2r2/go-i2c"
	"log"

	"github.com/d2r2/go-logger"
	"github.com/hamburgertrain/boostpi/internal/common"
	"github.com/hamburgertrain/boostpi/internal/configuration"
	"github.com/hamburgertrain/boostpi/internal/display"
	"github.com/hamburgertrain/boostpi/internal/elm327"
)

const (
	configFile string = "boostpi-config.json"
)

// Application entrypoint
func main() {
	log.Println("Loading configuration file:", configFile)
	var config configuration.Configuration
	config, err := configuration.LoadConfiguration(configFile)
	if err != nil {
		log.Fatalln("Error loading configuration:", err.Error())
	}
	log.Println("Configuration loaded")

	// Suppress/increase verbosity of output
	if !config.I2cDebug {
		err := logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
		if err != nil {
			log.Println("Could not set i2c log level:", err.Error())
		}
	}

	log.Println("Initializing connection to i2c display...")
	i2cDevice := display.Initialize(config)
	defer func(i2cDevice *i2c.I2C) {
		err := i2cDevice.Close()
		if err != nil {
			log.Println("Error closing i2c device:", err.Error())
		}
	}(i2cDevice)
	log.Println("Display connection initialized")

	// Let's make sure we have blank slate
	display.Reset(i2cDevice)

	// Show loading text while contacting ELM327 device
	display.ShowLoadingText(i2cDevice)

	// We might need to try this multiple times, device was not ready right away
	log.Println("Initializing connection to ELM327 device...")
	obdDevice, err := elm327.Initialize(config)
	if err != nil {
		display.ShowErrorAndShutdown(i2cDevice)
		log.Fatal("Could not initialize ELM327 device: ", err)
	}
	log.Println("ELM327 connection initialized")

	// Clear our display of loading text before showing boost
	display.Clear(i2cDevice)

	// Fetch information and display it
	common.GetAndDisplayValues(i2cDevice, obdDevice)

	log.Println("Turning display off")
	display.ShutdownDisplay(i2cDevice)
}
