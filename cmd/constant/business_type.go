package constant

import (
	"fmt"
	"strings"
)

type BusinessType int

const (
	Restaurant BusinessType = iota
	Cinema
	Other
	Religious
	Hospital
	All
)

func (businessType BusinessType) String() string {
	return [...]string{
		"Restaurant", // index 0, BusinessType = 0
		"Cinema",     // index 1, BusinessType = 1
		"Other",      // index 2, BusinessType = 2
		"Religious",  // index 3, BusinessType = 3
		"Hospital",   // index 4, BusinessType = 4
		"All",        // index 5, BusinessType = 5
	}[businessType] // Use businessType directly as an index
}

func ParseBusinessType(businessType string) (BusinessType, error) {
	switch strings.ToLower(businessType) {
	case "restaurant":
		return Restaurant, nil
	case "cinema":
		return Cinema, nil
	case "other":
		return Other, nil
	case "religious":
		return Religious, nil
	case "hospital":
		return Hospital, nil
	case "all":
		return All, nil
	default:
		return 0, fmt.Errorf("invalid business type")
	}
}
