package models

import (
	"errors"
	"fmt"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/utils"
)

//ContactOnClient is the data structure for persons contacts in clients.
type ContactOnClient struct {
	Name     string
	Phone    Telephone
	Whatsapp Telephone
	Email    string
	Skype    string
}

//Validate function validates if ContactOnClient object contains a minimum acceptable of data.
func (coc *ContactOnClient) Validate() error {

	if !utils.StringMatch(coc.Name, utils.PersonNameValidRegex) {
		return errors.New("Client contact name is invalid")
	}

	if e := coc.Phone.Validate(); e != nil {
		return fmt.Errorf("Client contact phone is invalid => %s", e.Error())
	}

	if e := coc.Whatsapp.Validate(); e != nil {
		return fmt.Errorf("Client contact whatsapp is invalid => %s", e.Error())
	}
	return nil
}
