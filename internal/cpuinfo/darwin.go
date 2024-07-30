package cpuinfo

import (
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	CTL_KERN    = 1
	KERN_CP_TIME = 34
	CTL_HW      = 6
	HW_NCPU     = 3
	KERN_PROC = 14
	KERN_PROC_ALL= 0
)

func getDarwinCPUInfo() (int, float64) {
	cores := getDarwinCPUCores()
	usages, err := getDarwinCPUUsage()
	if err != nil {
		panic(err)
	}
	return cores, usages
}

func getDarwinCPUCores() int {
	// Use runtime.NumCPU() as the primary method
	cores := runtime.NumCPU()

	// As a fallback, use sysctl command
	if cores == 0 {
		cmd := exec.Command("sysctl", "-n", "hw.ncpu")
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		cores, err = strconv.Atoi(strings.TrimSpace(string(output)))
		if err != nil {
			panic(err)
		}
	}

	return cores
}

func getDarwinCPUUsage() (float64, error) {
	cmd := exec.Command("top", "-l", "1", "-n", "0", "-S")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "CPU usage:") {
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				userCPU, err := strconv.ParseFloat(strings.TrimRight(fields[2], "%"), 64)
				if err != nil {
					return 0, err
				}
				sysCPU, err := strconv.ParseFloat(strings.TrimRight(fields[4], "%"), 64)
				if err != nil {
					return 0, err
				}
				return userCPU + sysCPU, nil
			}
		}
	}

	return 0, nil
}



