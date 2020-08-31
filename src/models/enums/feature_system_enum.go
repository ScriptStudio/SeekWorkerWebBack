package enums

import (
	"fmt"
)

//FeatureSystem is the features system enum.
type FeatureSystem uint16

const (
	//ChangeProfile is a FeatureSystem enum element that represents a system functionality
	ChangeProfile FeatureSystem = FeatureSystem(iota)

	//ClouseTasks is a FeatureSystem enum element that represents a system functionality
	ClouseTasks

	//ClouseCalls is a FeatureSystem enum element that represents a system functionality
	ClouseCalls

	//OpenTasks is a FeatureSystem enum element that represents a system functionality
	OpenTasks

	//OpenCalls is a FeatureSystem enum element that represents a system functionality
	OpenCalls

	//ReopenTasks is a FeatureSystem enum element that represents a system functionality
	ReopenTasks

	//ReopenCalls is a FeatureSystem enum element that represents a system functionality
	ReopenCalls
	// TODO: New options here.
)

//Validate checks whether the FeatureSystem is valid.
func (fs FeatureSystem) Validate() error {
	switch fs {
	case
		ChangeProfile,
		ClouseTasks,
		ClouseCalls,
		OpenTasks,
		OpenCalls,
		ReopenTasks,
		ReopenCalls:
		return nil
	}
	return fmt.Errorf("The FeatureSystem option is invalid: %d", fs)
}

//IsEqual checks whether two FeatureSystem is equals.
func (fs FeatureSystem) IsEqual(featSystem FeatureSystem) bool {
	return fs == featSystem
}
