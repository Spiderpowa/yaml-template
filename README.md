# YAML Templating Library for Golang

[![CircleCI](https://circleci.com/gh/Spiderpowa/yamltmpl.svg?style=shield)](https://circleci.com/gh/Spiderpowa/yamltmpl)

## Introduction

`yamltmpl` uses `text/template` as internal template engine. Check its [document](https://golang.org/pkg/text/template/) for details.

## Installation

```shell
go get github.com/Spiderpowa/yamltmpl
```

## Example

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/Spiderpowa/yamltmpl"
)

func main() {
    tmpl, err := yamltmpl.New("test", "host: {{.host}}\nport: {{.port}}\nname: it's a test")
    if err != nil {
        panic(err)
    }
    buf := new(bytes.Buffer)
    if err := tmpl.ApplyYaml([]byte("port: 3000\nhost: http://localhost"), buf); err != nil {
        panic(err)
    }
    fmt.Println(buf.String())
}
```

The output is

```yaml
host: http://localhost
port: 3000
name: it's a test
```

## Documentation

[Godoc](https://godoc.org/github.com/Spiderpowa/yamltmpl/template)
[Template](https://golang.org/pkg/text/template/)
