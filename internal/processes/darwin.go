package processes

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)


func getDarwinTopProcesses() ([]Process) {
	cmd := exec.Command("ps", "-Ao", "pid,comm,%cpu", "-r")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error executing ps command: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	var processes []Process

	// Start from index 1 to skip the header
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		pid, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		cpu, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			continue
		}

		processes = append(processes, Process{
			PID:  pid,
			Name: fields[1],
			CPU:  cpu,
		})
	}

	// Sort processes by CPU usage
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPU > processes[j].CPU
	})

	// Limit to top 10 processes
	if len(processes) > 10 {
		processes = processes[:10]
	}

	return processes
}