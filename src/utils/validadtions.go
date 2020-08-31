package utils

import "regexp"

const (
	_MaxLatitudeValue  float64 = 90
	_MinLatitudeValue  float64 = -90
	_MaxLongitudeValue float64 = 180
	_MinLongitudeValue float64 = -180

	//PersonNameValidRegex is the Person name Regex valid.
	PersonNameValidRegex string = `^((\b[A-zÀ-ú']{2,40}\b)\s*){2,}$`
)

//StringMatch validates a string based on a regex pattern.
//Return true if regex match sucess or false if not sucess
func StringMatch(value, pattern string) bool {
	var rgx *regexp.Regexp = regexp.MustCompile(pattern)

	return rgx.MatchString(value)
}

//CoordinatesValidation validates geographic coordinates.
func CoordinatesValidation(lat, long float64) (bool, float64) {
	if lat > _MaxLatitudeValue || lat < _MinLatitudeValue {
		return false, lat
	} else if long > _MaxLongitudeValue || long < _MinLongitudeValue {
		return false, long
	}
	return true, 0
}
