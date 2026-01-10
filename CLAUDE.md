# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

clinks is a minimal CLI tool for personal link profiles, written in Go. It reads a JSON configuration file containing a user's name and links, validates the data, and renders a formatted text output.

## Build and Test Commands

```bash
# Build the project
go build

# Run without building
go run main.go

# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests for a specific package
go test ./internal/profile

# Run a specific test
go test ./internal/profile -run TestProfileValidate

# Clean up dependencies
go mod tidy
```

## Architecture

The project follows a standard Go layout:

- `main.go` - Entry point, CLI handling, error formatting
- `internal/profile/` - Core package (not importable by external projects)
  - `profile.go` - Data structures (Profile, Link) and validation logic
  - `load.go` - Config file discovery and JSON loading
  - `render.go` - Text output rendering

### Key Design Decisions

- **URL Validation**: Go's `url.Parse()` is very permissive, so explicit validation checks for allowed schemes (http, https, mailto) and host presence for http/https URLs
- **Cross-Platform Paths**: Uses `os.UserConfigDir()` for platform-specific config paths rather than hardcoding `~/.config`
- **Package Naming**: The internal package uses `package profile` (not `package internal`) for clearer imports

### Config File Locations

clinks searches for `clinks.json` in order:
1. Current working directory
2. `$XDG_CONFIG_HOME/clinks/clinks.json` (typically `~/.config/clinks/clinks.json` on Linux)

### Config File Format

```json
{
  "name": "Your Name",
  "links": [
    {"label": "github", "url": "https://github.com/username"},
    {"label": "email", "url": "mailto:you@example.com"}
  ]
}
```
