package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Command int

const (
	CmdKill Command = iota
	CmdCheck
	CmdList
)

type model struct {
	command     Command
	port        int
	ports       []int
	stage       int
	animation   int
	killed      bool
	error       string
	processes   []ProcessInfo
	force       bool
	verbose     bool
	confirmed   bool
	showConfirm bool
}

type tickMsg time.Time
type processCheckMsg []ProcessInfo

func initialModel(cmd Command, port int, ports []int, force, verbose bool) model {
	return model{
		command:     cmd,
		port:        port,
		ports:       ports,
		stage:       0,
		force:       force,
		verbose:     verbose,
		showConfirm: cmd == CmdKill && !force,
	}
}

func (m model) Init() tea.Cmd {
	switch m.command {
	case CmdCheck:
		if len(m.ports) > 0 {
			return checkMultiplePorts(m.ports)
		}
		return checkSinglePort(m.port)
	case CmdList:
		return listAllPorts()
	default:
		return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			if m.command == CmdKill {
				if m.showConfirm && !m.confirmed {
					m.confirmed = true
					m.showConfirm = false
					m.stage = 1
					return m, tea.Tick(time.Millisecond*80, func(t time.Time) tea.Msg {
						return tickMsg(t)
					})
				} else if m.stage == 0 && m.force {
					m.stage = 1
					return m, tea.Tick(time.Millisecond*80, func(t time.Time) tea.Msg {
						return tickMsg(t)
					})
				}
			}
		case "n", "N":
			if m.showConfirm {
				return m, tea.Quit
			}
		}
	case tickMsg:
		if m.stage == 1 {
			m.animation++
			if m.animation > 25 {
				err := killProcessOnPort(m.port)
				if err != nil {
					m.error = err.Error()
					m.killed = false
				} else {
					m.killed = true
				}
				m.stage = 2
			}
			return m, tea.Tick(time.Millisecond*80, func(t time.Time) tea.Msg {
				return tickMsg(t)
			})
		}
	case processCheckMsg:
		m.processes = []ProcessInfo(msg)
		m.stage = 2
	}
	return m, nil
}

func (m model) View() string {
	// Dark hitman theme styles
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#DC143C")).
		Background(lipgloss.Color("#1C1C1C")).
		Padding(0, 1).
		MarginBottom(1)

	errorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF4444")).
		Bold(true)

	successStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")).
		Bold(true)

	infoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#CCCCCC"))

	warningStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true)

	switch m.command {
	case CmdCheck:
		return m.renderCheckView(titleStyle, infoStyle, errorStyle)
	case CmdList:
		return m.renderListView(titleStyle, infoStyle, errorStyle)
	default:
		return m.renderKillView(titleStyle, successStyle, errorStyle, warningStyle, infoStyle)
	}
}

func (m model) renderKillView(titleStyle, successStyle, errorStyle, warningStyle, infoStyle lipgloss.Style) string {
	var s string

	switch m.stage {
	case 0:
		if m.showConfirm {
			s += titleStyle.Render("⚠️  HITMAN - CONFIRMATION REQUIRED")
			s += "\n\n"
			s += warningStyle.Render(fmt.Sprintf("Target: Port %d", m.port))
			s += "\n\n"
			s += "Are you sure you want to eliminate this target?\n"
			s += infoStyle.Render("Press ENTER to confirm, 'n' to abort, 'q' to quit")
		} else {
			s += titleStyle.Render("🎯 HITMAN - READY TO ENGAGE")
			s += "\n\n"
			s += fmt.Sprintf("Target: Port %d\n", m.port)
			s += "\nPress ENTER to eliminate target\n"
			s += infoStyle.Render("Press 'q' to abort mission")
		}
	case 1:
		s += titleStyle.Render("🔫 HITMAN - ENGAGING TARGET")
		s += "\n\n"
		s += getAdvancedBulletAnimation(m.animation)
		s += "\n\n"
		s += fmt.Sprintf("Targeting port %d...\n", m.port)
		if m.verbose {
			s += infoStyle.Render(fmt.Sprintf("Animation frame: %d/25", m.animation))
		}
	case 2:
		if m.killed {
			s += successStyle.Render("✅ MISSION ACCOMPLISHED")
			s += "\n\n"
			s += "💀 Target eliminated successfully\n"
			s += fmt.Sprintf("Port %d is now free\n", m.port)
			if m.verbose {
				s += infoStyle.Render(fmt.Sprintf("Timestamp: %s", time.Now().Format("15:04:05")))
			}
		} else {
			s += errorStyle.Render("❌ MISSION FAILED")
			s += "\n\n"
			if m.error != "" {
				s += fmt.Sprintf("Error: %s\n", m.error)
			} else {
				s += "Target not found or already eliminated\n"
			}
		}
		s += "\n" + infoStyle.Render("Press 'q' to exit")
	}

	return s
}

func (m model) renderCheckView(titleStyle, infoStyle, errorStyle lipgloss.Style) string {
	var s string
	s += titleStyle.Render("🔍 HITMAN - RECONNAISSANCE")
	s += "\n\n"

	if m.stage < 2 {
		s += "Scanning targets...\n"
		s += getProgressAnimation()
	} else {
		if len(m.processes) == 0 {
			s += "No active targets found\n"
		} else {
			s += "Active targets identified:\n\n"
			for _, proc := range m.processes {
				s += fmt.Sprintf("Port %d: %s (PID %d)\n", proc.Port, proc.Name, proc.PID)
				if m.verbose {
					s += infoStyle.Render(fmt.Sprintf("  Command: %s\n", proc.Command))
				}
			}
		}
		s += "\n" + infoStyle.Render("Press 'q' to exit")
	}

	return s
}

func (m model) renderListView(titleStyle, infoStyle, errorStyle lipgloss.Style) string {
	var s string
	s += titleStyle.Render("📋 HITMAN - TARGET LIST")
	s += "\n\n"

	if m.stage < 2 {
		s += "Compiling target list...\n"
		s += getProgressAnimation()
	} else {
		if len(m.processes) == 0 {
			s += "No active targets found\n"
		} else {
			s += fmt.Sprintf("Found %d active targets:\n\n", len(m.processes))
			for _, proc := range m.processes {
				s += fmt.Sprintf("• Port %d: %s (PID %d)\n", proc.Port, proc.Name, proc.PID)
			}
		}
		s += "\n" + infoStyle.Render("Press 'q' to exit")
	}

	return s
}

func getAdvancedBulletAnimation(frame int) string {
	target := "🎯"

	// Different bullet stages
	var bullet string
	switch {
	case frame < 5:
		bullet = "🔫"
	case frame < 10:
		bullet = "💥"
	case frame < 15:
		bullet = "⚡"
	default:
		bullet = "💀"
	}

	width := 30
	pos := (frame * width) / 25

	animation := ""
	for i := 0; i < width; i++ {
		if i == pos {
			animation += bullet
		} else if i == width-1 {
			animation += target
		} else if i > pos && i < width-1 {
			animation += "═"
		} else {
			animation += " "
		}
	}

	return animation
}

func getProgressAnimation() string {
	dots := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	return dots[int(time.Now().UnixNano()/100000000)%len(dots)]
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		os.Exit(1)
	}

	args := os.Args[1:]
	var force, verbose bool
	var command Command = CmdKill
	var port int
	var ports []int

	// Parse flags
	filteredArgs := []string{}
	for _, arg := range args {
		switch arg {
		case "--force", "-f":
			force = true
		case "--verbose", "-v":
			verbose = true
		case "--help", "-h":
			showHelp()
			return
		case "--version":
			showVersion()
			return
		default:
			filteredArgs = append(filteredArgs, arg)
		}
	}

	if len(filteredArgs) == 0 {
		showHelp()
		os.Exit(1)
	}

	// Parse command
	switch filteredArgs[0] {
	case "check":
		command = CmdCheck
		if len(filteredArgs) < 2 {
			fmt.Println("❌ Check command requires a port or 'all'")
			os.Exit(1)
		}
		if filteredArgs[1] == "all" {
			command = CmdList
		} else if strings.Contains(filteredArgs[1], ",") {
			// Multiple ports
			portStrs := strings.Split(filteredArgs[1], ",")
			for _, portStr := range portStrs {
				p, err := strconv.Atoi(strings.TrimSpace(portStr))
				if err != nil || p < 1 || p > 65535 {
					fmt.Printf("❌ Invalid port: %s\n", portStr)
					os.Exit(1)
				}
				ports = append(ports, p)
			}
		} else {
			p, err := strconv.Atoi(filteredArgs[1])
			if err != nil || p < 1 || p > 65535 {
				fmt.Printf("❌ Invalid port: %s\n", filteredArgs[1])
				os.Exit(1)
			}
			port = p
		}
	case "list":
		command = CmdList
	default:
		// Kill command (default)
		p, err := strconv.Atoi(filteredArgs[0])
		if err != nil || p < 1 || p > 65535 {
			fmt.Printf("❌ Invalid port: %s\n", filteredArgs[0])
			os.Exit(1)
		}
		port = p
	}

	p := tea.NewProgram(initialModel(command, port, ports, force, verbose))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
