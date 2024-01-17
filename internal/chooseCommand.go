package internal

type SystemInfo interface{}

var err error

func GetSystemInfo(commandValue string) (SystemInfo, error) {
	var return_data SystemInfo

	switch commandValue {
	case "bios":
		return_data, err = BiosInfo()
	case "cpu":
		return_data, err = CpuInfo()
	case "gpu":
		return_data, err = GpuInfo()
	case "mem":
		return_data, err = MemoryInfo()
	case "network":
		return_data, err = NetworkInfo()
	}

	if err != nil {
		return nil, err
	}

	return return_data, nil
}
