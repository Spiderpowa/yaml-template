package yamltmpl

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

var defaultFunc = template.FuncMap{
	"base64":   b64,
	"readfile": readfile,
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
