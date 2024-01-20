// Facilitates connection to and execution of commands on an i2c LCD display
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
)

// i2c bus (0 -- original Pi, 1 -- Rev 2 Pi)
var i2cBus int = 1

// LCD Address
var i2cAddress uint8 = 0x27

// Commands
var lcdClearDisplay uint8 = 0x01
var lcdReturnHome uint8 = 0x02
var lcdEntryModeSet uint8 = 0x04
var lcdDisplayControl uint8 = 0x08
var lcdFunctionSet uint8 = 0x20

// Flags for display entry mode
var lcdEntryLeft uint8 = 0x02

// Flags for display on/off control
var lcdDisplayOn uint8 = 0x04

// Flags for function set
var lcd4BitMode uint8 = 0x00
var lcd2Line uint8 = 0x08
var lcd5x8Dots uint8 = 0x00

// Flags for backlight control
var lcdBacklight uint8 = 0x08
var lcdNoBacklight uint8 = 0x00

// Modes
var enableBit uint8 = 0b00000100         // Enable bit
var registerSelectBit uint8 = 0b00000001 // Register select bit

// Get our i2c connection
func Initialize() *i2c.I2C {
	connection, err := i2c.NewI2C(i2cAddress, i2cBus)
	if err != nil {
		log.Fatal("Could not initialize i2c device:", err)
	}

	return connection
}

// Put string function with optional char positioning
func LcdDisplayString(connection *i2c.I2C, str string, line uint8, pos uint8) {
	var posNew uint8 = 0

	if line == 1 {
		posNew = pos
	} else if line == 2 {
		posNew = 0x40 + pos
	} else if line == 3 {
		posNew = 0x14 + pos
	} else if line == 4 {
		posNew = 0x54 + pos
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
	time.Sleep(1 * time.Nanosecond)
}

// Clear lcd and set to home
func Clear(connection *i2c.I2C) {
	lcdWrite(connection, lcdClearDisplay, 0)
	lcdWrite(connection, lcdReturnHome, 0)
}

// Turn the display backlight off
func TurnBacklightOff(connection *i2c.I2C) {
	_ = writeCmd(connection, lcdNoBacklight)
}

// Write a single command
func writeCmd(connection *i2c.I2C, cmd uint8) int {
	// cast uint8 to byte in order to be written
	buf := make([]byte, 1)
	buf[0] = byte(cmd)

	res, err := connection.WriteBytes(buf)
	if err != nil {
		log.Fatal("Could not write to i2c device:", err)
	}
	time.Sleep(1 * time.Nanosecond)

	return res
}

// Clocks EN to latch command
func lcdStrobe(connection *i2c.I2C, data uint8) {
	writeCmd(connection, data|enableBit|lcdBacklight)
	time.Sleep(1 * time.Nanosecond)
	writeCmd(connection, ((data & ^enableBit) | lcdBacklight))
	time.Sleep(1 * time.Nanosecond)
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
