# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
