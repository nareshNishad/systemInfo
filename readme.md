# macOS CPU Info and Processes

This repository provides Go packages to retrieve CPU information and top processes on macOS.

## Packages

The project is divided into two packages:

1. `cpuinfo`: Provides functions to get CPU cores and usage information.
2. `processes`: Provides functions to retrieve the top processes by CPU usage.

## `cpuinfo` Package

### Functions

#### `getDarwinCPUInfo() (int, float64)`

Retrieves the number of CPU cores and the current CPU usage.

**Returns:**

- `int`: Number of CPU cores.
- `float64`: Total CPU usage (user + system).

#### `getDarwinCPUCores() int`

Retrieves the number of CPU cores using `runtime.NumCPU()` and as a fallback, the `sysctl` command.

**Returns:**

- `int`: Number of CPU cores.

#### `getDarwinCPUUsage() (float64, error)`

Retrieves the current CPU usage by executing the `top` command.

**Returns:**

- `float64`: Total CPU usage (user + system).
- `error`: Error encountered during execution, if any.

### Example Usage

```go
package main

import (
    "fmt"
    "cpuinfo"
)

func main() {
    cores, usage := cpuinfo.GetDarwinCPUInfo()
    fmt.Printf("CPU Cores: %d\n", cores)
    fmt.Printf("CPU Usage: %.2f%%\n", usage)
}
```
