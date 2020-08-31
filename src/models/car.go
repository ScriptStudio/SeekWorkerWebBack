package models

import (
	"errors"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/utils"
)

//CarPlateValidRgx is a regex validation of car plate.
const CarPlateValidRgx = "[A-Z]{3}[0-9]{4}|[A-Z]{3}[0-9][A-Z][0-9]{2}"

//Car is a struct for Car data
type Car struct {
	CarPlate     string
	TankCapacity float32
	Consumption  float32
}

//Validate function validates if Car object contains a minimum acceptable of data.
func (c *Car) Validate() error {

	if !utils.StringMatch(c.CarPlate, CarPlateValidRgx) {
		return errors.New("Car plate is invalid")
	}

	if c.TankCapacity <= 0 {
		return errors.New("Car tank capacity must be greater than zero")
	}

	if c.Consumption <= 0 {
		return errors.New("Car consumption must be greater than zero")
	}
	return nil
}

//Autonomy function returns the autonomy car.
func (c *Car) Autonomy() float32 {
	return c.TankCapacity * c.Consumption
}
