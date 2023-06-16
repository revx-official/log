# output

---

*output* is part of the *revx* project. It is a simple logging framework which enables colored console output.

Features:

- console logging
- colored console logging
- logging with different log levels

---

## Quickstart

```go
import (
  "github.com/revx-official/output/log"
)

func main() {
  log.Infof("Hello there!")
}
```

## Setup

To get *revx* up and running, follow the instructions below.

### Platforms

Officially supported development platforms are:

- Windows
- MacOS
- Linux

### Go

The *revx* project is written in *Go*, hence it is required to install *Go*. For the latest version of *Go*, check: https://go.dev/doc/install

## Usage

Using a fixed log level:

```go
import (
  "github.com/revx-official/output/log"
)

func init() {
  log.Level = log.LogLevelWarn
}

func main() {
  log.Infof("Info!")
  log.Warnf("Warning!")
}
```

Using flags:

```go
import (
  "github.com/revx-official/output/log"
)

func main() {
  var verbose bool

  flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
  flag.Parse()

  if verbose {
    log.Level = log.LogLevelTrace
  }

  log.Tracef("Trace!")
}
```

## Customization

To adjust the colors used for console logging, have a look at the following example:

```go
import (
  "github.com/revx-official/output/console"
)

func init() {
  console.ConsoleColorInfo = console.NewConsoleColor(FgColorGreen)
}

func main() {
  log.Infof("Hello there!")
}
```