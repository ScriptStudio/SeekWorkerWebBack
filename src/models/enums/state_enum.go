package enums

import "errors"

//State is a brazilian states enum.
type State uint8

const (
	//AC is a State enum element.
	AC State = 1 + iota
	//AL is a State enum element.
	AL
	//AP is a State enum element.
	AP
	//AM is a State enum element.
	AM
	//BA is a State enum element.
	BA
	//CE is a State enum element.
	CE
	//DF is a State enum element.
	DF
	//ES is a State enum element.
	ES
	//GO is a State enum element.
	GO
	//MA is a State enum element.
	MA
	//MT is a State enum element.
	MT
	//MS is a State enum element.
	MS
	//MG is a State enum element.
	MG
	//PA is a State enum element.
	PA
	//PB is a State enum element.
	PB
	//PR is a State enum element.
	PR
	//PE is a State enum element.
	PE
	//PI is a State enum element.
	PI
	//RJ is a State enum element.
	RJ
	//RN is a State enum element.
	RN
	//RS is a State enum element.
	RS
	//RO is a State enum element.
	RO
	//RR is a State enum element.
	RR
	//SC is a State enum element.
	SC
	//SP is a State enum element.
	SP
	//SE is a State enum element.
	SE
	//TO is a State enum element.
	TO
)

//Validate checks whether the State is valid.
func (s State) Validate() error {
	switch s {
	case
		AC, AL, AM, AP, BA, CE, DF, ES, GO, MA, MT, MS,
		MG, PA, PB, PR, PE, PI, RJ, RN, RS, RO, RR, SC,
		SP, SE, TO:
		return nil
	}
	return errors.New("The State option is invalid")
}
