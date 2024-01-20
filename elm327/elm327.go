// Facilitates connection to and execution of commands on an ELM327 OBD-II reader
package elm327

import (
	"log"

	"github.com/rzetterberg/elmobd"
)

var ELM327_DEVICE_LOCATION string = "/dev/ttyUSB0"
var ELM327_DEBUG bool = true

// Establish contact with an ELM327 OBD-II reader
func ContactElm327Device() {
	dev, err := elmobd.NewDevice(ELM327_DEVICE_LOCATION, ELM327_DEBUG)
	if err != nil {
		log.Println("Failed to create new device", err)
		return
	}

	version, err := dev.GetVersion()
	if err != nil {
		log.Println("Failed to get version", err)
		return
	}

	log.Println("Device has version", version)

	rpm, err := dev.RunOBDCommand(elmobd.NewEngineRPM())

	if err != nil {
		log.Println("Failed to get rpm", err)
		return
	}

	log.Printf("Engine spins at %s RPMs\n", rpm.ValueAsLit())
}
