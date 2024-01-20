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
var I2CBUS int = 1

// LCD Address
var ADDRESS uint8 = 0x27

// Commands
var LCD_CLEARDISPLAY uint8 = 0x01
var LCD_RETURNHOME uint8 = 0x02
var LCD_ENTRYMODESET uint8 = 0x04
var LCD_DISPLAYCONTROL uint8 = 0x08
var LCD_CURSORSHIFT uint8 = 0x10
var LCD_FUNCTIONSET uint8 = 0x20
var LCD_SETCGRAMADDR uint8 = 0x40
var LCD_SETDDRAMADDR uint8 = 0x80

// Flags for display entry mode
var LCD_ENTRYRIGHT uint8 = 0x00
var LCD_ENTRYLEFT uint8 = 0x02
var LCD_ENTRYSHIFTINCREMENT uint8 = 0x01
var LCD_ENTRYSHIFTDECREMENT uint8 = 0x00

// Flags for display on/off control
var LCD_DISPLAYON uint8 = 0x04
var LCD_DISPLAYOFF uint8 = 0x00
var LCD_CURSORON uint8 = 0x02
var LCD_CURSOROFF uint8 = 0x00
var LCD_BLINKON uint8 = 0x01
var LCD_BLINKOFF uint8 = 0x00

// Flags for display/cursor shift
var LCD_DISPLAYMOVE uint8 = 0x08
var LCD_CURSORMOVE uint8 = 0x00
var LCD_MOVERIGHT uint8 = 0x04
var LCD_MOVELEFT uint8 = 0x00

// Flags for function set
var LCD_8BITMODE uint8 = 0x10
var LCD_4BITMODE uint8 = 0x00
var LCD_2LINE uint8 = 0x08
var LCD_1LINE uint8 = 0x00
var LCD_5x10DOTS uint8 = 0x04
var LCD_5x8DOTS uint8 = 0x00

// Flags for backlight control
var LCD_BACKLIGHT uint8 = 0x08
var LCD_NOBACKLIGHT uint8 = 0x00

// Modes
var En uint8 = 0b00000100 // Enable bit
var Rw uint8 = 0b00000010 // Read/Write bit
var Rs uint8 = 0b00000001 // Register select bit

// Get our i2c connection
func Initialize() *i2c.I2C {
	connection, err := i2c.NewI2C(ADDRESS, I2CBUS)
	if err != nil {
		log.Fatal("Could not initialize i2c device:", err)
	}

	return connection
}

// Write a single command
func writeCmd(connection *i2c.I2C, cmd uint8) int {
	buf := make([]byte, 1)
	buf[0] = byte(cmd) // cast uint8 to byte

	res, err := connection.WriteBytes(buf)
	if err != nil {
		log.Fatal("Could not write to i2c device:", err)
	}
	time.Sleep(1 * time.Nanosecond)

	return res
}

// Reset display
func Reset(connection *i2c.I2C) {
	lcdWrite(connection, 0x03, 0)
	lcdWrite(connection, 0x03, 0)
	lcdWrite(connection, 0x03, 0)
	lcdWrite(connection, 0x02, 0)

	lcdWrite(connection, LCD_FUNCTIONSET|LCD_2LINE|LCD_5x8DOTS|LCD_4BITMODE, 0)
	lcdWrite(connection, LCD_DISPLAYCONTROL|LCD_DISPLAYON, 0)
	lcdWrite(connection, LCD_CLEARDISPLAY, 0)
	lcdWrite(connection, LCD_ENTRYMODESET|LCD_ENTRYLEFT, 0)
	time.Sleep(1 * time.Nanosecond)
}

// Clear lcd and set to home
func Clear(connection *i2c.I2C) {
	lcdWrite(connection, LCD_CLEARDISPLAY, 0)
	lcdWrite(connection, LCD_RETURNHOME, 0)
}

// Turn the display off
func TurnOff(connection *i2c.I2C) {
	_ = writeCmd(connection, LCD_NOBACKLIGHT)
}

// Clocks EN to latch command
func lcdStrobe(connection *i2c.I2C, data uint8) {
	writeCmd(connection, data|En|LCD_BACKLIGHT)
	time.Sleep(1 * time.Nanosecond)
	writeCmd(connection, ((data & ^En) | LCD_BACKLIGHT))
	time.Sleep(1 * time.Nanosecond)
}

// Write four bits
func lcdWriteFourBits(connection *i2c.I2C, data uint8) {
	writeCmd(connection, data|LCD_BACKLIGHT)
	lcdStrobe(connection, data)
}

// Write a command to lcd
func lcdWrite(connection *i2c.I2C, cmd uint8, mode uint8) {
	lcdWriteFourBits(connection, mode|(cmd&0xF0))
	lcdWriteFourBits(connection, mode|((cmd<<4)&0xF0))
}

// Write a character to lcd (or character rom) 0x09: backlight | RS=DR<
// works!
func lcdWriteChar(connection *i2c.I2C, charvalue uint8, mode uint8) {
	lcdWriteFourBits(connection, mode|(charvalue&0xF0))
	lcdWriteFourBits(connection, mode|((charvalue<<4)&0xF0))
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
		lcdWrite(connection, str[i], Rs)
	}
}
