package models

import (
	"errors"
	"fmt"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models/enums"
	"github.com/ScriptStudio/SeekWorkerWebBack/src/utils"
)

const (

	//ZipCodeValidRgx is a regex validation of zip codes.
	ZipCodeValidRgx string = "[0-9]{5}-[0-9]{3}"
)

//Address is a struct for Address data.
type Address struct {
	ZipCode      string
	State        enums.State
	City         string
	Neighborhood string
	Street       string
	Complement   string
	Coordinates  [2]float64 // [latitude, longitude]
}

//Validate function validates if Address object contains a minimum acceptable of data
func (addr *Address) Validate() error {
	if !utils.StringMatch(addr.ZipCode, ZipCodeValidRgx) {
		return errors.New("ZipCode is invalid")
	}
	if e := addr.State.Validate(); e != nil {
		return e
	}
	if len(addr.City) < 3 {
		return errors.New("City name is invalid")
	}
	if len(addr.Neighborhood) < 3 {
		return errors.New("Neighborhood name is invalid")
	}
	if len(addr.Street) < 3 {
		return errors.New("Street name is invalid")
	}
	if res, coord := utils.CoordinatesValidation(addr.Coordinates[0], addr.Coordinates[1]); !res {

		return fmt.Errorf("Addres coordinate invalid: %f", coord)
	}
	return nil
}
