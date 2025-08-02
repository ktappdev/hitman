package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ProcessInfo struct {
	PID     int
	Port    int
	Name    string
	Command string
}

// killProcessOnPort kills the process running on the specified port
func killProcessOnPort(port int) error {
	pid, err := findProcessOnPort(port)
	if err != nil {
		return err
	}

	if pid == 0 {
		return fmt.Errorf("no process found on port %d", port)
	}

	return killProcess(pid)
}

// findProcessOnPort finds the PID of the process using the specified port
func findProcessOnPort(port int) (int, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin", "linux":
		cmd = exec.Command("lsof", "-ti", fmt.Sprintf(":%d", port))
	case "windows":
		cmd = exec.Command("netstat", "-ano")
	default:
		return 0, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to find process on port %d", port)
	}

	if runtime.GOOS == "windows" {
		return parseWindowsNetstat(string(output), port)
	}

	pidStr := strings.TrimSpace(string(output))
	if pidStr == "" {
		return 0, nil
	}

	pids := strings.Split(pidStr, "\n")
	pid, err := strconv.Atoi(strings.TrimSpace(pids[0]))
	if err != nil {
		return 0, fmt.Errorf("invalid PID: %s", pids[0])
	}

	return pid, nil
}

// getProcessInfo gets detailed information about a process on a port
func getProcessInfo(port int) (*ProcessInfo, error) {
	pid, err := findProcessOnPort(port)
	if err != nil || pid == 0 {
		return nil, err
	}

	name, command := getProcessDetails(pid)

	return &ProcessInfo{
		PID:     pid,
		Port:    port,
		Name:    name,
		Command: command,
	}, nil
}

// getProcessDetails gets the name and command of a process by PID
func getProcessDetails(pid int) (string, string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin", "linux":
		cmd = exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "comm=,args=")
	case "windows":
		cmd = exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid), "/FO", "CSV", "/NH")
	default:
		return "unknown", "unknown"
	}

	output, err := cmd.Output()
	if err != nil {
		return "unknown", "unknown"
	}

	if runtime.GOOS == "windows" {
		return parseWindowsTasklist(string(output))
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) > 0 {
		parts := strings.SplitN(strings.TrimSpace(lines[0]), " ", 2)
		if len(parts) >= 2 {
			return parts[0], parts[1]
		} else if len(parts) == 1 {
			return parts[0], parts[0]
		}
	}

	return "unknown", "unknown"
}

// parseWindowsTasklist parses Windows tasklist output
func parseWindowsTasklist(output string) (string, string) {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) > 0 {
		// CSV format: "name","pid","session","mem"
		parts := strings.Split(lines[0], ",")
		if len(parts) > 0 {
			name := strings.Trim(parts[0], "\"")
			return name, name
		}
	}
	return "unknown", "unknown"
}

// parseWindowsNetstat parses netstat output to find PID for the port
func parseWindowsNetstat(output string, port int) (int, error) {
	lines := strings.Split(output, "\n")
	portStr := fmt.Sprintf(":%d ", port)

	for _, line := range lines {
		if strings.Contains(line, portStr) && strings.Contains(line, "LISTENING") {
			fields := strings.Fields(line)
			if len(fields) >= 5 {
				pid, err := strconv.Atoi(fields[len(fields)-1])
				if err == nil {
					return pid, nil
				}
			}
		}
	}

	return 0, nil
}

// killProcess kills the process with the specified PID
func killProcess(pid int) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin", "linux":
		cmd = exec.Command("kill", "-9", strconv.Itoa(pid))
	case "windows":
		cmd = exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Run()
}

// getAllListeningPorts gets all ports with listening processes
func getAllListeningPorts() ([]ProcessInfo, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin", "linux":
		cmd = exec.Command("lsof", "-i", "-P", "-n", "-sTCP:LISTEN")
	case "windows":
		cmd = exec.Command("netstat", "-ano", "-p", "TCP")
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	if runtime.GOOS == "windows" {
		return parseWindowsNetstatAll(string(output))
	}

	return parseUnixLsofAll(string(output))
}

// parseUnixLsofAll parses lsof output for all listening ports
func parseUnixLsofAll(output string) ([]ProcessInfo, error) {
	var processes []ProcessInfo
	lines := strings.Split(output, "\n")

	for _, line := range lines[1:] { // Skip header
		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 9 {
			pidStr := fields[1]
			addressPort := fields[8]

			pid, err := strconv.Atoi(pidStr)
			if err != nil {
				continue
			}

			// Extract port from address:port
			parts := strings.Split(addressPort, ":")
			if len(parts) < 2 {
				continue
			}

			port, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				continue
			}

			name, command := getProcessDetails(pid)

			processes = append(processes, ProcessInfo{
				PID:     pid,
				Port:    port,
				Name:    name,
				Command: command,
			})
		}
	}

	return processes, nil
}

// parseWindowsNetstatAll parses Windows netstat output for all listening ports
func parseWindowsNetstatAll(output string) ([]ProcessInfo, error) {
	var processes []ProcessInfo
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if !strings.Contains(line, "LISTENING") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 5 {
			addressPort := fields[1]
			pidStr := fields[len(fields)-1]

			pid, err := strconv.Atoi(pidStr)
			if err != nil {
				continue
			}

			parts := strings.Split(addressPort, ":")
			if len(parts) < 2 {
				continue
			}

			port, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				continue
			}

			name, command := getProcessDetails(pid)

			processes = append(processes, ProcessInfo{
				PID:     pid,
				Port:    port,
				Name:    name,
				Command: command,
			})
		}
	}

	return processes, nil
}

// Tea commands for async operations
func checkSinglePort(port int) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(500 * time.Millisecond) // Simulate work
		proc, err := getProcessInfo(port)
		if err != nil || proc == nil {
			return processCheckMsg([]ProcessInfo{})
		}
		return processCheckMsg([]ProcessInfo{*proc})
	}
}

func checkMultiplePorts(ports []int) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(800 * time.Millisecond) // Simulate work
		var processes []ProcessInfo

		for _, port := range ports {
			proc, err := getProcessInfo(port)
			if err == nil && proc != nil {
				processes = append(processes, *proc)
			}
		}

		return processCheckMsg(processes)
	}
}

func listAllPorts() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1000 * time.Millisecond) // Simulate work
		processes, err := getAllListeningPorts()
		if err != nil {
			return processCheckMsg([]ProcessInfo{})
		}
		return processCheckMsg(processes)
	}
}
