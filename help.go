package main

import (
	"fmt"
)

func showHelp() {
	help := `
🎯 HITMAN - Elite Process Terminator

DESCRIPTION:
    A dark, stylish terminal-based process killer with advanced reconnaissance
    capabilities. Eliminates processes running on specified ports with smooth
    animations and a gritty hitman-themed UI.

USAGE:
    hitman <port>                    Kill process on specified port
    hitman check <port>              Check what's running on a port
    hitman check <port1,port2,...>   Check multiple ports
    hitman check all                 Check all occupied ports
    hitman list                      List all occupied ports
    hitman -h, --help               Show this help message
    hitman --version                Show version information

FLAGS:
    --force, -f                     Skip confirmation prompts
    --verbose, -v                   Show detailed output and timestamps

EXAMPLES:
    hitman 3000                     Kill process on port 3000
    hitman check 8080               Check what's running on port 8080
    hitman check 3000,8080,9000     Check multiple ports
    hitman check all                Scan all occupied ports
    hitman list                     List all processes with ports
    hitman 3000 --force             Kill without confirmation
    hitman check 8080 --verbose     Detailed process information

CONTROLS:
    ENTER / SPACE                   Execute action/confirm
    n / N                          Decline confirmation
    q / Ctrl+C                     Quit/Abort mission

FEATURES:
    🔫 Smooth bullet animations during elimination
    🎯 Dark hitman-themed UI with red/black colors
    🔍 Advanced process reconnaissance
    ⚡ Cross-platform compatibility
    📋 Detailed process information
    ⚠️  Confirmation prompts for safety
    🎨 Professional CLI experience

SUPPORTED PLATFORMS:
    • macOS (Darwin) - Uses lsof, ps, kill
    • Linux - Uses lsof, ps, kill  
    • Windows - Uses netstat, tasklist, taskkill

REQUIREMENTS:
    • macOS/Linux: lsof, ps, kill (usually pre-installed)
    • Windows: netstat, tasklist, taskkill (built-in)

Made with ❤️  for developers who like their process killing with style.
`
	fmt.Print(help)
}

func showVersion() {
	fmt.Println("🎯 HITMAN v2.0.0 - Elite Edition")
	fmt.Println("The ultimate stylish process terminator")
	fmt.Println("Dark theme • Advanced reconnaissance • Smooth animations")
}
