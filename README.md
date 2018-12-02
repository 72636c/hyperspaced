# H Y P E R S P A C E D

[![Build Status](https://cloud.drone.io/api/badges/72636c/hyperspaced/status.svg)](https://cloud.drone.io/72636c/hyperspaced)

Command line utilities to improve the aesthetics of your favourite phrases with
automatic space insertion (ASI).

## Usage

### CLI

Prerequisites:

- Go 1.11+

```shell
go install ./cmd/spaced
```

```shell
echo 'gofmt urself' | spaced
```

### Go

```go
import (
  "github.com/72636c/hyperspaced"
)

hyperspaced.Spaced("gofmt urself")
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
