# 🎯 HITMAN - Elite Process Terminator

> *"Precision. Style. Elimination."*

A dark, stylish terminal-based CLI tool with advanced reconnaissance capabilities. Eliminates processes running on specified ports with smooth animations and a gritty hitman-themed UI built with Go, Bubble Tea, and Lipgloss.

![Version](https://img.shields.io/badge/version-2.0.0-red?style=flat-square)
![Go](https://img.shields.io/badge/go-1.21+-blue?style=flat-square)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-green?style=flat-square)
![License](https://img.shields.io/badge/license-MIT-yellow?style=flat-square)

## ✨ Features

- � **Prmecision Elimination** - Kill processes on any port with style and safety confirmations
- �️ **Advanced Reconnaissance** - Intel gathering on what's running on specific ports
- � **Tarbget Enumeration** - Scan and list all occupied ports on your system
- 💥 **Cinematic Animations** - Multi-frame bullet sequences during elimination missions
- 🎨 **Dark Operative Theme** - Professional red/black UI with smooth transitions
- 🖥️ **Cross-Platform Operations** - Seamless execution on macOS, Linux, and Windows
- ⚡ **Lightning Fast** - Optimized process detection and termination
- 📖 **Mission Briefings** - Comprehensive help system with tactical examples
- 🛡️ **Safety Protocols** - Confirmation prompts with --force override capability
- 🔧 **Intel Mode** - Verbose output for detailed process information and timestamps

## 🚀 Deployment

### Quick Install from Source
```bash
git clone https://github.com/ktappdev/hitman.git
cd hitman
go mod tidy
go build -o hitman
```

### System-wide Installation via Go
```bash
# Install directly from GitHub (requires Go 1.21+)
go install github.com/ktappdev/hitman@latest

# Or install from a specific version/tag
go install github.com/ktappdev/hitman@v2.0.0
```

### Manual Build
```bash
# Clone and build manually
git clone https://github.com/ktappdev/hitman.git
cd hitman
go mod download
go build -ldflags="-s -w" -o hitman .

# Make executable and move to PATH (optional)
chmod +x hitman
sudo mv hitman /usr/local/bin/
```

### Binary Release
Download pre-compiled binaries from [Releases](https://github.com/ktappdev/hitman/releases)

## 📖 Mission Manual

### 🎯 Elimination Operations
```bash
# Standard elimination (with confirmation protocol)
./hitman 3000

# Silent takedown (bypass confirmation)
./hitman 8080 --force

# Detailed elimination report
./hitman 9000 --verbose --force
```

### 🔍 Reconnaissance Missions
```bash
# Single target reconnaissance
./hitman check 3000

# Multi-target surveillance
./hitman check 3000,8080,9000,5432

# Full network scan
./hitman check all

# Complete target enumeration
./hitman list

# Deep intel gathering
./hitman check 8080 --verbose
```

### 📋 Command Center
```bash
# Mission briefing
./hitman --help

# System status
./hitman --version
```

## 🎮 Operator Controls

| Input | Action |
|-------|--------|
| `ENTER` / `SPACE` | Execute mission / Confirm operation |
| `n` / `N` | Abort confirmation / Decline mission |
| `q` / `Ctrl+C` | Emergency extraction / Quit |

## 🎯 Mission Types

| Command | Objective | Example |
|---------|-----------|---------|
| `<port>` | Eliminate target on specified port | `hitman 3000` |
| `check <port>` | Gather intel on single target | `hitman check 8080` |
| `check <ports>` | Multi-target reconnaissance | `hitman check 3000,8080` |
| `check all` | Network-wide surveillance | `hitman check all` |
| `list` | Complete target enumeration | `hitman list` |

## 🚩 Tactical Flags

| Flag | Alias | Mission Enhancement |
|------|-------|-------------------|
| `--force` | `-f` | Bypass all confirmation protocols |
| `--verbose` | `-v` | Enable detailed mission reporting |
| `--help` | `-h` | Access mission briefing |
| `--version` | | Display system specifications |

## ⚙️ Operational Mechanics

### 🔍 Target Acquisition
- **macOS/Linux**: Utilizes `lsof` for precision port-to-PID mapping
- **Windows**: Deploys `netstat` for comprehensive port scanning
- **Process Intel**: Leverages `ps` (Unix) / `tasklist` (Windows) for target profiling

### 💥 Elimination Sequence
Multi-stage visual feedback during termination:
```
🔫 → 💥 → ⚡ → 💀
```
- Smooth cross-terminal bullet trajectory
- Professional status indicators
- Real-time mission progress

### 🎯 Termination Protocols
- **Unix Systems**: `kill -9` for immediate termination
- **Windows**: `taskkill /F` for forced process elimination

## 🖥️ Platform Compatibility Matrix

| Platform | Detection | Profiling | Termination | Status |
|----------|-----------|-----------|-------------|---------|
| **macOS** | `lsof` | `ps` | `kill -9` | ✅ Fully Operational |
| **Linux** | `lsof` | `ps` | `kill -9` | ✅ Fully Operational |
| **Windows** | `netstat` | `tasklist` | `taskkill /F` | ✅ Fully Operational |

## 📋 System Requirements

### Runtime Dependencies
- **Unix Systems**: `lsof`, `ps`, `kill` (standard system tools)
- **Windows**: `netstat`, `tasklist`, `taskkill` (built-in utilities)

### Build Requirements
- **Go**: 1.21+ for compilation
- **Dependencies**: Automatically managed via `go mod`

## 🎨 Visual Experience

Experience the dark art of process termination with:

- **🎭 Elite Operative Theme**: Professional red/black color palette
- **🎬 Cinematic Animations**: Smooth bullet trajectories and transitions  
- **📊 Real-time Feedback**: Progress indicators and status updates
- **🛡️ Safety Protocols**: Clear confirmation prompts and error handling
- **📱 Responsive Design**: Adapts to various terminal sizes

## 🔧 Mission Examples

### Quick Elimination
```bash
# Terminate development server
./hitman 3000
```

### Stealth Operations
```bash
# Silent elimination of multiple targets
./hitman 8080 --force
./hitman 9000 --force
```

### Intelligence Gathering
```bash
# Comprehensive network reconnaissance
./hitman check 3000,8080,9000,5432 --verbose

# Full system surveillance
./hitman list --verbose
```

### Development Workflow
```bash
# Check if your React dev server is running
./hitman check 3000

# Eliminate stuck Node.js processes
./hitman 8080 --force

# Scan for common development ports
./hitman check 3000,8080,9000,5000,4000
```

## 🤝 Contributing

Contributions welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) before submitting missions.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎖️ Acknowledgments

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for terminal UI
- Styled with [Lipgloss](https://github.com/charmbracelet/lipgloss) for visual excellence
- Inspired by the need for stylish process management

---

<div align="center">

**Made with ❤️ for developers who demand precision, style, and efficiency in their terminal operations.**

*"In the world of process management, there are no second chances. Choose HITMAN."*

</div># hitman
