package enums

import "errors"

//EmployeeType is a system user profile enum.
type EmployeeType uint8

const (
	//Outside is a EmployeeType enum element that represents outside workers.
	Outside EmployeeType = 1 + iota

	//Retaguard is a EmployeeType enum element that represents system operator.
	Retaguard

	//Admin is a EmployeeType enum element that represents system admin.
	Admin
)

//Validate checks whether the EmployeeType is valid.
func (et EmployeeType) Validate() error {
	switch et {
	case
		Outside,
		Retaguard,
		Admin:
		return nil
	}
	return errors.New("EmployeeType option is invalid")
}
