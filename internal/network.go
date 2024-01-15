package internal

import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/net"
)

func NetworkInfo() (*net.Info, error) {
	network, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}
	return network, err
}
