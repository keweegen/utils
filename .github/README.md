<h1 align="center">Golang utils</h1>
<p align="center">
  <a href="https://goreportcard.com/report/github.com/keweegen/utils">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B?style=flat-square">
  </a>
  <a href="https://gocover.io/github.com/keweegen/utils">
    <img src="https://img.shields.io/badge/%F0%9F%94%8E%20gocover-97.8%25-75C46B.svg?style=flat-square">
  </a>
  <a href="https://github.com/keweegen/utils/actions?query=workflow%3ASecurity">
    <img src="https://img.shields.io/github/workflow/status/keweegen/utils/Security?label=%F0%9F%94%91%20gosec&style=flat-square&color=75C46B">
  </a>
  <a href="https://github.com/keweegen/utils/actions?query=workflow%3ATest">
    <img src="https://img.shields.io/github/workflow/status/keweegen/utils/Test?label=%F0%9F%A7%AA%20tests&style=flat-square&color=75C46B">
  </a>
</p>

## ⚙ Installation

```shell
go get -u github.com/keweegen/utils
```

## 👀 Examples

#### Errors

**Code**

```go
const (
    ErrApp = iota + 1
    ErrReadConfig
)

func main() {
    fmt.Printf("-- Default settings --\n")
    printErrors()

    // -
    
    errors.CurrentSettings.
        SetDefaultCode(-1).
        SetSeparator(" -> ").
        SetErrorFormatter(func(err errors.KError) string {
            return fmt.Sprintf("(%d: %s)", err.Code(), err.Message())
        })

    fmt.Printf("\n-- Custom settings [DefaultCode, Separator, ErrorFormatter] --\n")
    printErrors()
}

func printErrors() {
    errReadConfig := errors.New(ErrReadConfig, "read config")
    errApp := errors.New(ErrApp, "init app", errReadConfig)
    
    fmt.Printf("Output   = %s\n", errApp)
    fmt.Printf("Unwrap   = %s\n", errors.Unwrap(errApp))
    fmt.Printf("Wrap     = %s\n", errors.Wrap(errApp, 0, "wrap #1"))
    fmt.Printf("Wrap fmt = %s\n", fmt.Errorf("wrap #2: %w", errApp))
}
```

**Output**

```
-- Default settings --
Output   = [ERR1] init app: [ERR2] read config
Unwrap   = [ERR2] read config
Wrap     = [ERR0] wrap #1: [ERR1] init app: [ERR2] read config
Wrap fmt = wrap #2: [ERR1] init app: [ERR2] read config

-- Custom settings [DefaultCode, Separator, ErrorFormatter] --
Output   = (1: init app) -> (2: read config)
Unwrap   = (2: read config)
Wrap     = (-1: wrap #1) -> (1: init app) -> (2: read config)
Wrap fmt = wrap #2: (1: init app) -> (2: read config)
```
