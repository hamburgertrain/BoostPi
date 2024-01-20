// Facilitates connection to and execution of commands on an ELM327 OBD-II reader
package elm327

import (
	"log"

	"github.com/rzetterberg/elmobd"
)

var ELM327_DEVICE_LOCATION string = "/dev/ttyUSB0"
var ELM327_DEBUG bool = true

// Establish contact with an ELM327 OBD-II reader nd read some stats
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

	GetEngineRpm(dev)
	GetIntakeManifoldPressure(dev)
	GetMassAirflowRate(dev)
}

// Get engine rpm from OBD device
func GetEngineRpm(dev *elmobd.Device) {
	rpm, err := dev.RunOBDCommand(elmobd.NewEngineRPM())
	if err != nil {
		log.Println("Failed to get rpm", err)
		return
	}

	log.Printf("Engine spinning at %s RPMs\n", rpm.ValueAsLit())
}

// Get intake manifold pressure from OBD device
func GetIntakeManifoldPressure(dev *elmobd.Device) {
	imaf, err := dev.RunOBDCommand(elmobd.NewIntakeManifoldPressure())
	if err != nil {
		log.Println("Failed to get imaf", err)
		return
	}

	log.Printf("Intake Manifold Pressure is %s\n", imaf.ValueAsLit())
}

// Get mass airflow rate from OBD device
func GetMassAirflowRate(dev *elmobd.Device) {
	mafr, err := dev.RunOBDCommand(elmobd.NewMafAirFlowRate())
	if err != nil {
		log.Println("Failed to get mafr", err)
		return
	}

	log.Printf("Mass airflow rate is %s\n", mafr.ValueAsLit())
}
