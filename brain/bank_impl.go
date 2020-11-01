package brain

import (
	"fmt"
	"github.com/gansidui/ahocorasick"
)

func (bb *bankBrain) pickupProperties(msg string) (bool, propertiesVec) {
	validVal := propertiesVec{}
	fmt.Println("aaa")
	return true, validVal
}

func (bb *bankBrain) pickupName(msg string) (bool, string) {
	ac := ahocorasick.NewMatcher()

	ac.Build(bb.allNames)

	return true, ""
}

func (bb *bankBrain) pickupMobilePhone(msg string) (bool, string) {
	return true, ""
}
func (bb *bankBrain) pickupDomain(msg string) (bool, string) {
	return true, ""
}
