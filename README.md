# H Y P E R S P A C E D

[![Build Status](https://cloud.drone.io/api/badges/72636c/hyperspaced/status.svg)](https://cloud.drone.io/72636c/hyperspaced)
[![Latest Release](https://img.shields.io/github/release/72636c/hyperspaced.svg?logo=github)](https://github.com/72636c/hyperspaced/releases/latest)
[![GoDoc](https://godoc.org/github.com/72636c/hyperspaced?status.svg)](https://godoc.org/github.com/72636c/hyperspaced)
[![Go Report Card](https://goreportcard.com/badge/github.com/72636c/hyperspaced)](https://goreportcard.com/report/github.com/72636c/hyperspaced)

Utilities to improve the aesthetics of your favourite phrases with automatic
space insertion (ASI).

## Usage

### CLI

```shell
$ echo 'AESTHETIC' | spaced
A E S T H E T I C

$ echo 'AESTHETIC' | spaced 2 lower
a  e  s  t  h  e  t  i  c
```

### Go

```go
import (
  "github.com/72636c/hyperspaced"
)

hyperspaced.One("AESTHETIC")
// A E S T H E T I C

hyperspaced.LowerN("AESTHETIC", 2)
// a  e  s  t  h  e  t  i  c
```

## Meta

### CI/CD pipeline

Builds and releases are automated with [Drone](https://drone.io/), which is
impressively _"powered by blazing fast bare-metal servers"_, and more
importantly _"written in Go"_.

### Local scripts

```shell
./scripts/build-package.sh
```

```shell
./scripts/lint-test.sh
```

### Mouseprint

_This is a shameless rip-off of <https://www.npmjs.com/package/letter-spacing>._
