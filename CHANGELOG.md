# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- added following commands:
  - Mixer:
    - KEYER
    - CHROMA
    - BLEND
    - INVERT
    - OPACITY
    - BRIGHTNESS

### Changed

- changed URLs in ReadMe to point to correct package

## [0.1.0] - 24-03-2026

### Removed

- removed direct CasparCG Server return and wrapped error code in custom error object

### Changed

- split up `builder.go` into multiple subfiles

### Added

- custom error object for CasparCG server returns
- `.vscode` settings
- `.golangci.yaml` linting and formatting
- `.github/workflows` enforcement of updating changelog
- added examples explanation to `ReadMe.md`

### Updated

- `.github/workflows` to use new linter

### Fixed

- check changelog running on main
- a bunch of linter and formatter issues

## [0.0.1] 28-02-2026

### Added

- logo under `./assets/` that merges the go-gopher and casparCG logo
- examples under `./cmd/`
- generic types
- DTOs for most returns
- following commands:
  - Generic:
    - LOAD
    - PLAY
    - PAUSE
    - RESUME
    - STOP
    - CLEAR
    - CALL
    - SWAP
    - ADD
    - REMOVE
    - PRINT
    - LOGLEVEL
    - LOCK
    - PING
  - Template:
    - CG ADD
    - CG PLAY
    - CG STOP
    - CG NEXT
    - CG REMOVE
    - CG CLEAR
    - CG UPDATE
    - CG INVOKE
    - CG INFO
  - Query:
    - CINF
    - CLS
    - FLS
    - TLS
    - VERSION
    - INFO
    - INFO TEMPLATE
    - INFO CONFIG
    - INFO PATHS
    - INFO SYSTEM
    - INFO SERVER
    - INFO QUEUES
    - INFO THREADS
    - INFO DELAY
    - BYE
    - KILL
    - RESTART
