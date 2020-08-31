package models

import (
	"errors"
	"fmt"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models/enums"
	"github.com/ScriptStudio/SeekWorkerWebBack/src/utils"
)

const (

	//ImeiValidRegex is a regex validation of smartphone imei.
	ImeiValidRegex string = "[0-9]{15}"
)

//Employee is a struct for Employee data.
type Employee struct {
	Name          string
	Phone         Telephone
	Email         string
	Addr          Address
	Photo         string
	Password      string
	Imei          string
	Coordinates   [2]float64
	Profile       Profile
	EmpType       [3]enums.EmployeeType
	ActuationArea [3]float64 // [latitude, longitude, radius(in meters)]
}

//Validate function validates if Employee object contains a minimum acceptable of data.
func (ep *Employee) Validate() error {

	if !utils.StringMatch(ep.Name, utils.PersonNameValidRegex) {
		return errors.New("The employee name is invalid")
	}

	if e := ep.Phone.Validate(); e != nil {
		return e
	}

	if e := ep.Addr.Validate(); e != nil {
		return e
	}

	if !utils.StringMatch(ep.Imei, ImeiValidRegex) {
		return errors.New("The IMEI number is invalid")
	}

	if res, coord := utils.CoordinatesValidation(ep.Coordinates[0], ep.Coordinates[1]); !res {
		return fmt.Errorf("Employee coordinate invalid: %f", coord)
	}

	if e := ep.Profile.Validate(); e != nil {
		return e
	}

	if len(ep.EmpType) > 0 {
		for _, empType := range ep.EmpType {
			if e := empType.Validate(); e != nil {
				return e
			}
		}
	} else {
		return errors.New("The employee type list cannot be empty")
	}

	if len(ep.ActuationArea) > 0 {
		if res, coord := utils.CoordinatesValidation(ep.ActuationArea[0], ep.ActuationArea[1]); !res {
			return fmt.Errorf("Employee actuation coordinate invalid: %f", coord)
		}

		if ep.ActuationArea[2] <= 0 {
			return errors.New("The employee actuation area radius must be greater than zero")
		}
	}
	return nil
}
