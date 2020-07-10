package models

import (
	"fmt"
	"strconv"
)

const (

	//TelephoneDddValueMin is the minimum value of Telephone ddd valid.
	TelephoneDddValueMin uint8 = 11

	//TelephoneDddValueMax is the maximum value of Telephone ddd valid.
	TelephoneDddValueMax uint8 = 99
)

//Telephone is a struct for telephone data
type Telephone struct {
	Ddd    uint8
	Number uint32
}

//FormatedTelephone returns a full and formated Telephone as a string
func (t Telephone) FormatedTelephone() string {
	var ddd, number string
	ddd = strconv.Itoa(int(t.Ddd))
	number = strconv.Itoa(int(t.Number))

	return fmt.Sprintf("(%s)%s", ddd, number)
}

//CleanedTelephone returns a full and cleaned Telephone as a string
func (t Telephone) CleanedTelephone() string {
	var ddd, number string = strconv.Itoa(int(t.Ddd)), strconv.Itoa(int(t.Number))
	return ddd + number
}
