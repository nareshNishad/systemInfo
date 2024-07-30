package cpuinfo

import (
	"runtime"
)

// GetCPUInfo retrieves the number of CPU cores and the CPU usage percentage
func GetCPUInfo() (int, float64) {
	switch runtime.GOOS {
	case "darwin":
		return getDarwinCPUInfo()
	default:
		return 0, 0.0
	}
}

