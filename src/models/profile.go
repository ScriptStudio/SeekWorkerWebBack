package models

import (
	"errors"
	"fmt"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models/enums"
)

//Profile is a structure for user profile data.
type Profile struct {
	Name        string
	Permissions map[enums.FeatureSystem]bool
	Description string
}

//Validate function validates if Profile object contains a minimum acceptable of data.
func (p *Profile) Validate() error {

	if len(p.Permissions) > 0 {
		for permiss := range p.Permissions {
			if e := permiss.Validate(); e != nil {
				return fmt.Errorf("The profile permissions list is invalid => %s", e.Error())
			}
		}
	} else {
		return errors.New("The permissions list cannot be empty")
	}
	return nil
}

//SetOnePermission is the Profile function that set a permission to a ProfileActions
func (p *Profile) SetOnePermission(fs enums.FeatureSystem, permission bool) error {

	if e := fs.Validate(); e != nil {
		return e
	}

	p.Permissions[fs] = permission
	return nil
}

//GetOnePermission is the Profile function that get a especific permission
func (p *Profile) GetOnePermission(fs enums.FeatureSystem) (bool, error) {

	if _, present := p.Permissions[fs]; !present {
		return false, errors.New("Feature permission non-exists or is invalid")
	}
	return p.Permissions[fs], nil
}
