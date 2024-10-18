// Houses functions related to loading configuration files
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
package configuration

import (
	"encoding/json"
	"log"
	"os"
)

// Representation of the contents of our `boostpi-config.json`
type Configuration struct {
	I2cBus               int    // Which i2c bus are we using? (0 for older pi revisions, 1 for newer)
	I2cAddress           string // The address of the i2c device on the bus
	I2cDebug             bool   // Should i2c be running in debug or info
	Elm327DeviceLocation string // USB device location
	Elm327Debug          bool   // elm327 debug messaging
}

// Loads a JSON configuration file into our Configuration struct
func LoadConfiguration(configFileName string) (Configuration, error) {
	file, err := os.Open(configFileName)
	if err != nil {
		return Configuration{}, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing configuration file:", err.Error())
		}
	}(file)

	decoder := json.NewDecoder(file)

	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}
