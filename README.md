# YAML Templating Library for Golang

[![CircleCI](https://circleci.com/gh/Spiderpowa/yaml-template.svg?style=shield)](https://circleci.com/gh/Spiderpowa/yaml-template)

## Introduction

`yaml-template` uses `text/template` as internal template engine. Check its [document](https://golang.org/pkg/text/template/) for details.

## Installation

```shell
go get github.com/Spiderpowa/yaml-template
```

## Example

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/Spiderpowa/yaml-template/template"
)

func main() {
    tmpl, err := template.New("test", "host: {{.host}}\nport: {{.port}}\nname: it's a test")
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

[Godoc](https://godoc.org/github.com/Spiderpowa/yaml-template/template)
[Template](https://golang.org/pkg/text/template/)
