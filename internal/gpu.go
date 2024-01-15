package internal

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func GpuInfo() (*ghw.GPUInfo, error) {
	gpu, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting gpu info: %v", err)
	}
	return gpu, err
}
