# H Y P E R S P A C E D

Command line utilities to improve the aesthetics of your favourite phrases with
automatic space insertion (ASI).

## CLI

```shell
go install ./cmd/spaced
```

```shell
echo 'gofmt urself' | spaced
```

## Go

```go
import (
  "github.com/72636c/hyperspaced"
)

hyperspaced.Spaced("gofmt urself")
```

## Mouseprint

_This is a shameless clone of <https://www.npmjs.com/package/letter-spacing>._
