package models

import (
	"errors"
	"fmt"
	"strconv"
)

const (

	//TelephoneDddValueMin is the minimum value of Telephone ddd valid.
	TelephoneDddValueMin uint8 = 11

	//TelephoneDddValueMax is the maximum value of Telephone ddd valid.
	TelephoneDddValueMax uint8 = 99

	//TelephoneNumberMinLen is the mininun value of Telephone number valid.
	TelephoneNumberMinLen uint8 = 8

	//TelephoneNumberMaxLen is the mininun value of Telephone number valid.
	TelephoneNumberMaxLen uint8 = 9
)

//Telephone is a struct for telephone data
type Telephone struct {
	Ddd    uint8
	Number uint32
}

//Validate function validates if Telephone object contains a minimum acceptable of data.
func (t *Telephone) Validate() error {

	lenNumber := uint8(len(strconv.Itoa(int(t.Number))))

	if t.Ddd < TelephoneDddValueMin || t.Ddd > TelephoneDddValueMax {
		return errors.New("Telephone DDD is invalid")
	}
	if lenNumber < TelephoneNumberMinLen || lenNumber > TelephoneNumberMaxLen {
		return errors.New("Telephone number is invalid")
	}
	return nil
}

//FormatedTelephone returns a full and formated Telephone as a string
func (t *Telephone) FormatedTelephone() string {

	var ddd, number string

	ddd = strconv.Itoa(int(t.Ddd))
	number = strconv.Itoa(int(t.Number))

	numlen := len(number)

	return fmt.Sprintf("(%s) %s-%s", ddd, number[:numlen-4], number[numlen-4:])
}

//CleanedTelephone returns a full and cleaned Telephone as a string
func (t *Telephone) CleanedTelephone() string {
	var ddd, number string = strconv.Itoa(int(t.Ddd)), strconv.Itoa(int(t.Number))
	return ddd + number
}
