package tests

import (
	"testing"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models"
)

const errorMsn = "Expected value %s, Obtained value %s"

func TestMethods(t *testing.T) {

	phone := models.Telephone{Ddd: 49, Number: 988070350}

	var testFormatedPhone = func(t *testing.T) {
		expecValue := "(49) 98807-0350"
		var obtainedValue string = phone.FormatedTelephone()

		if expecValue != obtainedValue {
			t.Errorf(errorMsn, expecValue, obtainedValue)
		}
	}

	var testCleanedTelephone = func(t *testing.T) {
		expecValue := "49988070350"
		var obtainedValue string = phone.CleanedTelephone()

		if expecValue != obtainedValue {
			t.Errorf(errorMsn, expecValue, obtainedValue)
		}
	}

	t.Run("FormatedTelephoneMethod", testFormatedPhone)
	t.Run("CleanedTelephoneMethod", testCleanedTelephone)
}

func TestElements(t *testing.T) {

}
