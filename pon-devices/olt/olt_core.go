package devices

import (
	"errors"
	"log"
	"os"
)

type Olt struct {
	Name  string        `json:"Name"`
	Onus  map[int8]*Onu `json:"Onu-Map"`
	Npon  int8          `json:"num-of-pon-ports"`
	Nonus int8          `json:num-of-onus"`
}

func printBanner() {
	text := `

	   	   ___  _   _____
          / _ \| | |_   _|
         | | | | |   | |
         | |_| | |___| |
          \___/|_____|_|`
	log.Println(text)
}

func newoltdevice(name string, nonus int8, npon int8) (*Olt, error) {
	NewOlt := &Olt{Name: name,
		Npon:  npon,
		Nonus: nonus}
	if nonus >= 127 {
		log.Println("No-of-onus greater than the limit")
		return errors.New("No-of-onus greater than the limit")

	}
	return NewOlt, nil
}

func StartOltDevice(name string, nonus int8, npon int8) {
	log.Println("Info: Starting olt device")
	olt, err := newoltdevice(name, nonus, npon)
	if err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
