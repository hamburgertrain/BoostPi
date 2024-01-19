// (eventually) Pull from ELM327 device and write boost data to an i2c LCD
package main

import (
	"fmt"
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

func main() {
	fmt.Println("Initializing connection...")
	// Create new connection to I2C bus on 2 line with address 0x27
	i2cConnection, err := i2c.NewI2C(ADDRESS, I2CBUS)
	if err != nil {
		log.Fatal(err)
	}
	// Free I2C connection on exit
	defer i2cConnection.Close()
	fmt.Println("Connection initialized")

	fmt.Println("Resetting display...")
	Reset(i2cConnection)
	time.Sleep(1 * time.Second)
	fmt.Println("Display reset")

	fmt.Println("Writing to display...")
	LcdDisplayString(i2cConnection, "Hello world!", 1, 0)
	fmt.Println("Writing done.")

	time.Sleep(5 * time.Second)

	fmt.Println("Clearing display...")
	Clear(i2cConnection)
	fmt.Println("Clear")

	time.Sleep(1 * time.Second)

	for i := 0; i < 30; i++ {
		LcdDisplayString(i2cConnection, string(i), 1, 0)
		time.Sleep(1 * time.Second)
		Clear(i2cConnection)
		time.Sleep(2 * time.Second)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Turning display off")
	_ = WriteCmd(i2cConnection, LCD_NOBACKLIGHT)
}

// Write a single command
func WriteCmd(connection *i2c.I2C, cmd uint8) int {
	fmt.Printf("Value came in as: %d\n", cmd)
	buf := make([]byte, 1)
	buf[0] = byte(cmd) // cast uint8 to byte
	fmt.Printf("Writing value as: %d\n", buf[0])
	res, err := connection.WriteBytes(buf)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// Reset display
func Reset(connection *i2c.I2C) {
	LcdWrite(connection, 0x03, 0)
	LcdWrite(connection, 0x03, 0)
	LcdWrite(connection, 0x03, 0)
	LcdWrite(connection, 0x02, 0)

	LcdWrite(connection, LCD_FUNCTIONSET|LCD_2LINE|LCD_5x8DOTS|LCD_4BITMODE, 0)
	LcdWrite(connection, LCD_DISPLAYCONTROL|LCD_DISPLAYON, 0)
	LcdWrite(connection, LCD_CLEARDISPLAY, 0)
	LcdWrite(connection, LCD_ENTRYMODESET|LCD_ENTRYLEFT, 0)
}

// Clear lcd and set to home
func Clear(connection *i2c.I2C) {
	LcdWrite(connection, LCD_CLEARDISPLAY, 0)
	LcdWrite(connection, LCD_RETURNHOME, 0)
}

// Clocks EN to latch command
func LcdStrobe(connection *i2c.I2C, data uint8) {
	WriteCmd(connection, data|En|LCD_BACKLIGHT)
	time.Sleep(1)
	WriteCmd(connection, ((data & ^En) | LCD_BACKLIGHT))
	time.Sleep(1)
}

// Write four bits I guess
func LcdWriteFourBits(connection *i2c.I2C, data uint8) {
	WriteCmd(connection, data|LCD_BACKLIGHT)
	LcdStrobe(connection, data)
}

// Write a command to lcd
func LcdWrite(connection *i2c.I2C, cmd uint8, mode uint8) {
	LcdWriteFourBits(connection, mode|(cmd&0xF0))
	LcdWriteFourBits(connection, mode|((cmd<<4)&0xF0))
}

// Write a character to lcd (or character rom) 0x09: backlight | RS=DR<
// works!
func LcdWriteChar(connection *i2c.I2C, charvalue uint8, mode uint8) {
	LcdWriteFourBits(connection, mode|(charvalue&0xF0))
	LcdWriteFourBits(connection, mode|((charvalue<<4)&0xF0))
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

	LcdWrite(connection, 0x80+posNew, 0)

	for i := 0; i < len(str); i++ {
		fmt.Printf("Writing %s\n", str[i])
		LcdWrite(connection, str[i], Rs)
	}

}
