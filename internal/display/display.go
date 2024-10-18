// Facilitates connection to and execution of commands on an i2c LCD 16x2 display
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
package display

import (
	"log"
	"time"

	"github.com/d2r2/go-i2c"
	"github.com/hamburgertrain/boostpi/internal/configuration"
	"github.com/hamburgertrain/boostpi/internal/utilities"
)

const (
	// Commands
	lcdClearDisplay   uint8 = 0x01
	lcdReturnHome     uint8 = 0x02
	lcdEntryModeSet   uint8 = 0x04
	lcdDisplayControl uint8 = 0x08
	lcdFunctionSet    uint8 = 0x20

	// Flags for display entry mode
	lcdEntryLeft uint8 = 0x02

	// Flags for display on/off control
	lcdDisplayOn  uint8 = 0x04
	lcdDisplayOff uint8 = 0x00

	// Flags for function set
	lcd4BitMode uint8 = 0x00
	lcd2Line    uint8 = 0x08
	lcd5x8Dots  uint8 = 0x00

	// Flags for backlight control
	lcdBacklight   uint8 = 0x08
	lcdNoBacklight uint8 = 0x00

	// Modes
	enableBit         uint8 = 0b00000100 // Enable bit
	registerSelectBit uint8 = 0b00000001 // Register select bit

	// Display text
	loadingTextLine1 string = "----BoostPi-----"
	loadingTextLine2 string = "----Loading-----"
	errorTextLine1   string = "ERROR"
	errorTextLine2   string = "SHUTTING DOWN"
)

// Get our i2c connection
func Initialize(config configuration.Configuration) *i2c.I2C {
	i2cAddress, err := utilities.ConvertToUint8(config.I2cAddress)
	if err != nil {
		log.Fatal("Could not parse i2cAddress from config:", err.Error())
	}

	connection, err := i2c.NewI2C(i2cAddress, config.I2cBus)
	if err != nil {
		log.Fatal("Could not initialize i2c device:", err.Error())
	}

	return connection
}

// Show our loading text
func ShowLoadingText(connection *i2c.I2C) {
	LcdDisplayString(connection, loadingTextLine1, 1, 0)
	LcdDisplayString(connection, loadingTextLine2, 2, 0)
}

// Clear display, show error text and shutdown display
func ShowErrorAndShutdown(connection *i2c.I2C) {
	Clear(connection)
	showError(connection)
	time.Sleep(5 * time.Second)
	ShutdownDisplay(connection)
}

// Put string function with char positioning
func LcdDisplayString(connection *i2c.I2C, str string, line uint8, pos uint8) {
	var posNew uint8 = 0

	// First or second row/line?
	if line == 1 {
		posNew = pos
	} else if line == 2 {
		posNew = 0x40 + pos
	}

	lcdWrite(connection, 0x80+posNew, 0)

	for i := 0; i < len(str); i++ {
		lcdWrite(connection, str[i], registerSelectBit)
	}
}

// Reset display
func Reset(connection *i2c.I2C) {
	lcdWrite(connection, 0x03, 0)
	lcdWrite(connection, 0x03, 0)
	lcdWrite(connection, 0x03, 0)
	lcdWrite(connection, 0x02, 0)

	lcdWrite(connection, lcdFunctionSet|lcd2Line|lcd5x8Dots|lcd4BitMode, 0)
	lcdWrite(connection, lcdDisplayControl|lcdDisplayOn, 0)
	lcdWrite(connection, lcdClearDisplay, 0)
	lcdWrite(connection, lcdEntryModeSet|lcdEntryLeft, 0)
	time.Sleep(200 * time.Millisecond)
}

// Clear lcd and set to home
func Clear(connection *i2c.I2C) {
	lcdWrite(connection, lcdClearDisplay, 0)
	lcdWrite(connection, lcdReturnHome, 0)
}

// Turn off backlight and display after clearing
func ShutdownDisplay(connection *i2c.I2C) {
	Clear(connection)
	turnBacklightOff(connection)
	turnDisplayOff(connection)
}

// Show our error text
func showError(connection *i2c.I2C) {
	LcdDisplayString(connection, errorTextLine1, 1, 0)
	LcdDisplayString(connection, errorTextLine2, 2, 0)
}

// Turn the backlight off
func turnBacklightOff(connection *i2c.I2C) {
	writeCmd(connection, lcdNoBacklight)
}

// Turn the display off
func turnDisplayOff(connection *i2c.I2C) {
	writeCmd(connection, lcdDisplayOff)
}

// Write a single command
func writeCmd(connection *i2c.I2C, cmd uint8) {
	// cast uint8 to byte in order to be written
	buf := make([]byte, 1)
	buf[0] = cmd

	_, err := connection.WriteBytes(buf)
	if err != nil {
		log.Println("Could not write to i2c device:", err.Error())
	}
	time.Sleep(100 * time.Microsecond)
}

// Clocks EN to latch command
func lcdStrobe(connection *i2c.I2C, data uint8) {
	writeCmd(connection, data|enableBit|lcdBacklight)
	time.Sleep(500 * time.Microsecond)
	writeCmd(connection, (data & ^enableBit)|lcdBacklight)
	time.Sleep(100 * time.Microsecond)
}

// Write four bits
func lcdWriteFourBits(connection *i2c.I2C, data uint8) {
	writeCmd(connection, data|lcdBacklight)
	lcdStrobe(connection, data)
}

// Write a command to lcd
func lcdWrite(connection *i2c.I2C, cmd uint8, mode uint8) {
	lcdWriteFourBits(connection, mode|(cmd&0xF0))
	lcdWriteFourBits(connection, mode|((cmd<<4)&0xF0))
}
