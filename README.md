<p align="center">
  <img src="./assets/gophor_caspar.png" width="250" alt="CasparCG AMCP Go Logo">
</p>

<h1 align="center">CasparCG AMCP Go</h1>

<p align="center">
  <b>A fast, idiomatic Go implementation of the CasparCG AMCP protocol.</b>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/OverlayFox/casparcg-amcp-go"><img src="https://pkg.go.dev/badge/github.com/OverlayFox/casparcg-amcp-go.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/OverlayFox/casparcg-amcp-go"><img src="https://goreportcard.com/badge/github.com/OverlayFox/casparcg-amcp-go" alt="Go Report Card"></a>
  <a href="https://casparcg.com/"><img src="https://img.shields.io/badge/CasparCG-AMCP-blue.svg" alt="CasparCG AMCP"></a>
</p>

---

## 📖 Overview

`casparcg-amcp-go` provides a clean, concurrent interface to communicate with [CasparCG](https://casparcg.com/) servers using [Go](https://go.dev/). It abstracts away the raw socket communication, handling connection management, command parsing, and response mapping so you can focus on building your broadcast logic.

Commands are built directly based on the official [AMCP Protocol Documentation](https://github.com/CasparCG/help/wiki/AMCP-Protocol).

### ✨ Features

- **Typed Responses:** Parses raw CasparCG returns into easy-to-use Data Transfer Objects (DTOs).
- **Expanded Commands:** Extends the base AMCP commands for easier execution and strict typing.
- **Optional Arguments:** Handles optional AMCP parameters using Go pointers.

> **⚠️ Note:** Not all AMCP commands have been implemented yet.
> The CasparCG Server version used to develop this package was `2.4.3`. Due to some inconsistencies in its returns, certain DTOs might be incomplete.

---

## 🚀 Getting Started

### Installation

Install the package via `go get`:

```bash
go get github.com/OverlayFox/casparcg-amcp-go
```

---

## 🚧 Roadmap & Missing Features

As it stands right now, you can use this package to **Play, Stop, and Resume Videos and Templates**, as well as get general info about the CasparCG Server and Media Server.

The main feature categories currently missing are:

- Batching
- Data Commands
- Mixer Commands

Below is a specific list of AMCP commands still to be implemented:

- [ ] [LOG CATEGORY](https://github.com/CasparCG/help/wiki/AMCP-Protocol#log-category)
- [ ] [SET](https://github.com/CasparCG/help/wiki/AMCP-Protocol#set)
- [ ] [All DATA Commands](https://github.com/CasparCG/help/wiki/AMCP-Protocol#data-commands)
- [ ] [All MIXER Commands](https://github.com/CasparCG/help/wiki/AMCP-Protocol#mixer-commands)
- [ ] [All Thumbnail Commands](https://github.com/CasparCG/help/wiki/AMCP-Protocol#thumbnail-commands)
- [ ] [DIAG](https://github.com/CasparCG/help/wiki/AMCP-Protocol#diag)
- [ ] [GL INFO](https://github.com/CasparCG/help/wiki/AMCP-Protocol#gl-info)
- [ ] [GL GC](https://github.com/CasparCG/help/wiki/AMCP-Protocol#gl-gc)
- [ ] [HELP](https://github.com/CasparCG/help/wiki/AMCP-Protocol#help)
- [ ] [HELP PRODUCER](https://github.com/CasparCG/help/wiki/AMCP-Protocol#help-producer)
- [ ] [HELP CONSUMER](https://github.com/CasparCG/help/wiki/AMCP-Protocol#help-consumer)
- [ ] [All Batching Commands](https://github.com/CasparCG/help/wiki/AMCP-Protocol#batching-commands)
