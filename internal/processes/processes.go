package processes

import (
	"runtime"
)

type Process struct {
	PID  int
	Name string
	CPU  float64
}

func GetTopProcesses() []Process {
	switch runtime.GOOS {
	case "darwin":
		return getDarwinTopProcesses()
	default:
		return nil
	}
}
