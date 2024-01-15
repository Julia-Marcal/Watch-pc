package internal

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func BiosInfo() (*ghw.BIOSInfo, error) {
	bios, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("Error getting BIOS info: %v", err)
	}
	return bios, err
}
