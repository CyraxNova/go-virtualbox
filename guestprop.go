package virtualbox

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var (
	getRegexp = regexp.MustCompile("^Value: (.*)$")
)

// SetGuestProperty writes a VirtualBox guestproperty to the given value.
func SetGuestProperty(vm string, prop string, val string) error {
	return vbm("guestproperty", "set", vm, prop, val)
}

// GetGuestProperty reads a VirtualBox guestproperty.
func GetGuestProperty(vm string, prop string) (string, error) {
	var out string
	var err error
	out, err = vbmOut("guestproperty", "get", vm, prop)
	if err != nil {
		log.Print(err)
		return "", err
	}
	out = strings.TrimSpace(out)
	if Verbose {
		log.Printf("out (trimmed): '%s'", out)
	}
	var match = getRegexp.FindStringSubmatch(out)
	if Verbose {
		log.Print("match:", match)
	}
	if len(match) != 2 {
		return "", fmt.Errorf("No match with VBoxManage get guestproperty output")
	} else {
		return match[1], nil
	}
}

// DeleteGuestProperty deletes a VirtualBox guestproperty.
func DeleteGuestProperty(vm string, prop string) error {
	return vbm("guestproperty", "delete", vm, prop)
}
