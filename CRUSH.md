# CRUSH

Build
- go mod tidy
- go build -v ./...
- go build -ldflags="-s -w" -o hitman .

Test
- go test ./...
- go test -v ./...
- Single test file: go test -v ./path/to/pkg -run TestName
- With race: go test -race ./...

Lint/Format/Static analysis
- go fmt ./...
- go vet ./...
- golangci-lint run

Release
- goreleaser release --clean

Coding style (Go)
- Imports: standard lib first, then external; use goimports/gofmt; keep aliases minimal
- Formatting: enforced by gofmt; no custom linters required beyond vet/golangci-lint
- Types: prefer explicit types; small structs; export only when needed
- Naming: CamelCase for exported, lowerCamelCase for internal; keep package names short, lowercased, no underscores
- Errors: return wrapped errors with context using fmt.Errorf("…: %w", err); no panics in lib paths; user-facing errors printed via fmt.Println in main
- CLI: keep argument parsing in main; showHelp() and showVersion() for UX; exit non-zero on invalid input
- Concurrency: use tea.Cmd for async UI tasks; avoid blocking UI; keep sleeps only for simulated work
- OS ops: use exec.Command for lsof/netstat/ps/taskkill; guard by runtime.GOOS and return clear errors for unsupported OS
- Logging/UI: Bubble Tea for TUI flow; Lipgloss for styling; no logs of sensitive data

Repo conventions
- Go version: 1.21+ (see go.mod)
- CI runs: go test ./..., go vet ./..., gofmt check, build matrix 1.21/1.22 and OSes
- Binaries/artifacts go to dist/ via GoReleaser; ldflags set version, commit, date

Editor/AI assistant rules
- If using Cursor/Copilot rules: none found in repo; follow above style
