<p align="center">
  <img src="./assets/gophor_caspar.png" width="250" alt="CasparCG AMCP Go Logo">
</p>

<h1 align="center">CasparCG AMCP Go</h1>

<p align="center">
  <b>A fast, idiomatic Go implementation of the CasparCG AMCP protocol.</b>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/overlayfox/casparcg-amcp-go"><img src="https://pkg.go.dev/badge/github.com/overlayfox/casparcg-amcp-go.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/overlayfox/casparcg-amcp-go"><img src="https://goreportcard.com/badge/github.com/overlayfox/casparcg-amcp-go" alt="Go Report Card"></a>
  <a href="https://casparcg.com/"><img src="https://img.shields.io/badge/CasparCG-AMCP-blue.svg" alt="CasparCG AMCP"></a>
  <a href="https://github.com/OverlayFox/casparcg-amcp-go/actions/workflows/go.yml"><img src="https://github.com/OverlayFox/casparcg-amcp-go/actions/workflows/go.yml/badge.svg" alt="Build Status"></a>
  <a href="https://go.dev/dl/"><img src="https://img.shields.io/badge/Go-%3E%3D1.21-blue?logo=go" alt="Minimum Go Version"></a>
</p>

---

## 📖 Overview

`casparcg-amcp-go` provides a clean, concurrent interface to communicate with [CasparCG](https://casparcg.com/) servers using [Go](https://go.dev/). It abstracts away the raw socket communication, handling command parsing and response mapping so you can focus on building your broadcast logic.

Commands are built directly based on the official [AMCP Protocol Documentation](https://casparcg.com/docs/wiki/protocols/amcp-protocol) with a more strict approach to optional variables.

## ✨ Features

- **Fluent Builder API:** Inspired by patterns like `zerolog`, commands are constructed through a chainable, context-aware builder.
- **Semantic AMCP Extensions:** The standard AMCP command were expanded into a more descriptive API. This replaces undocumented server strings with human-readable methods.
- **Rich Error Handling:** Includes a native CasparCGError struct that wraps raw server codes and messages into a structured format.
- **Explicit Parameter Safety:** Eliminates fallbacks to defaults `{[layer:int]|-0}`, by enforcing strict definitions for optional variables.
- **DTOs:** Commands return a structured DTO instead of raw data.

---

## 🧩 Examples

```golang
err := client.CG().Channel(1).Layer(1).CGLayer(1).Add(types.CGAdd{Template: "L3", PlayOnLoad: true, Data: jsonString})
if err != nil {
  var casparErr casparcg.CasparCGError
  if errors.As(err, &casparErr) {
    fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
  }
  panic(err)
}
```

More examples can be found in the [`./cmd/`](./cmd/) directory to see how to use this package. <br>
These are more starting points than full implementations.

## 🚀 Getting Started

### Installation

Install the package via `go get`:

```bash
go get github.com/OverlayFox/casparcg-amcp-go
```

---

## 🔬 Tests

This package enforces a strict approach to stability via a automated testing suite:

- **Strict Code Standards:** Enforces consistent code style using `.golangci.yaml` via multiple linters and formatters.
- [![Build Check & Unittest](https://github.com/OverlayFox/casparcg-amcp-go/actions/workflows/go.yml/badge.svg)](https://github.com/OverlayFox/casparcg-amcp-go/actions/workflows/go.yml) validates and tests every commit against all major Go versions via automated runners to ensure backward compatibility.
- ~~[![Integration Test](https://github.com/OverlayFox/casparcg-amcp-go/actions/workflows/casparTest.yml/badge.svg)](https://github.com/OverlayFox/casparcg-amcp-go/actions/workflows/casparTest.yml) verifies package against the latest builds of CasparCG Server `2.4` and `2.5`.~~ Not implemented yet.

## 🚧 Roadmap & Missing Features

As it stands right now, most of the commands have been implemented. <br>
But there are some more abstract commands that have not made it into this package.

The main feature categories currently missing are:

- Batching
- Data Commands

Below is a specific list of AMCP commands still to be implemented:

- [ ] [LOG CATEGORY](https://casparcg.com/docs/wiki/protocols/amcp-protocol#log-category)
- [ ] [All DATA Commands](https://casparcg.com/docs/wiki/protocols/amcp-protocol#data-commands)
- [ ] [All Thumbnail Commands](https://casparcg.com/docs/wiki/protocols/amcp-protocol#thumbnail-commands)
- [ ] [All Batching Commands](https://casparcg.com/docs/wiki/protocols/amcp-protocol#batching-commands)
- [ ] [MIXER DEFER Command](https://casparcg.com/docs/wiki/protocols/amcp-protocol#mixer-commit)
