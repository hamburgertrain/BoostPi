// Facilitates connection to and execution of commands on an ELM327 OBD-II reader
package elm327

import (
	"log"

	"github.com/rzetterberg/elmobd"
)

var ELM327_DEVICE_LOCATION string = "/dev/ttyUSB0"
var ELM327_DEBUG bool = true

// Establish contact with an ELM327 OBD-II reader
func Initialize() (*elmobd.Device, error) {
	dev, err := elmobd.NewDevice(ELM327_DEVICE_LOCATION, ELM327_DEBUG)
	if err != nil {
		log.Println("Failed to create new device", err)
		return nil, err
	}

	return dev, nil
}

// Get version from an ELM327 OBD-II reader
func GetVersion(dev *elmobd.Device) {
	version, err := dev.GetVersion()
	if err != nil {
		log.Println("Failed to get version", err)
		return
	}

	log.Println("Device has version", version)
}

// Get engine rpm from an ELM327 OBD-II reader
func GetEngineRpm(dev *elmobd.Device) {
	rpm, err := dev.RunOBDCommand(elmobd.NewEngineRPM())
	if err != nil {
		log.Println("Failed to get rpm", err)
		return
	}

	log.Printf("Engine spinning at %s RPMs\n", rpm.ValueAsLit())
}

// Get intake manifold pressure from an ELM327 OBD-II reader
func GetIntakeManifoldPressure(dev *elmobd.Device) {
	imaf, err := dev.RunOBDCommand(elmobd.NewIntakeManifoldPressure())
	if err != nil {
		log.Println("Failed to get imaf", err)
		return
	}

	log.Printf("Intake Manifold Pressure is %s\n", imaf.ValueAsLit())
}

// Get mass airflow rate from an ELM327 OBD-II reader
func GetMassAirflowRate(dev *elmobd.Device) {
	mafr, err := dev.RunOBDCommand(elmobd.NewMafAirFlowRate())
	if err != nil {
		log.Println("Failed to get mafr", err)
		return
	}

	log.Printf("Mass airflow rate is %s\n", mafr.ValueAsLit())
}
