package yamltmpl

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var defaultFunc = template.FuncMap{
	"base64":       b64,
	"readfile":     readfile,
	"castToString": toString,
	"indentSpace":  indentSpace,
	"indentTab":    indentTab,
	"indentWith":   indentWith,
}

func b64(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}

func readfile(file string) ([]byte, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	content, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}
	return content, nil
}

func toString(bytes []byte) string {
	return string(bytes)
}

func indentSpace(input string, indent int) string {
	return indentWith(input, indent, " ")
}

func indentTab(input string, indent int) string {
	return indentWith(input, indent, "\t")
}

func indentWith(input string, indent int, with string) string {
	indention := ""
	for i := 0; i < indent; i++ {
		indention += with
	}
	return strings.ReplaceAll(input, "\n", "\n"+indention)
}

// Funcs adds the default function.
func Funcs(funcs template.FuncMap) {
	for k, v := range funcs {
		defaultFunc[k] = v
	}
}
