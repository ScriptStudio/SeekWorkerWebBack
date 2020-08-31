package models

import (
	"fmt"
)

//Client is the structure for Client data.
type Client struct {
	Name         string
	Phone        Telephone
	Email        string
	Addr         Address
	Logo         string
	Contacts     []ContactOnClient
	Observations string
}

//Validate function validates if Client object contains a minimum acceptable of data.
func (c *Client) Validate() error {

	if e := c.Phone.Validate(); e != nil {
		return fmt.Errorf("Client phone is invalid => %s", e.Error())
	}

	if e := c.Addr.Validate(); e != nil {
		return fmt.Errorf("Client address is invalid => %s", e.Error())
	}

	for _, contac := range c.Contacts {
		if e := contac.Validate(); e != nil {
			return fmt.Errorf("The client contacts list is invalid => %s", e.Error())
		}
	}
	return nil
}
