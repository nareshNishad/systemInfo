package main

import (
	"cpuinfo/internal/cpuinfo"
	"cpuinfo/internal/processes"
	"fmt"
)

func main() {
	fmt.Println("Detecting CPU Info and Top Processes...")

	cpuCores, cpuUsage := cpuinfo.GetCPUInfo()
	fmt.Printf("Number of CPU cores: %d\n", cpuCores)
	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)

	topProcesses := processes.GetTopProcesses()
	fmt.Println("Top 5 processes by CPU usage:")
	for i, proc := range topProcesses {
		fmt.Printf("%d: PID %d (%s) - %.2f%%\n", i+1, proc.PID, proc.Name, proc.CPU)
	}
}
