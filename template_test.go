package yamltmpl

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile(".", "test_template.yaml")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("host: {{.host}}\nport: {{.port}}\nname: it's a test")
	tmpFile.Close()

	tmpl := New("test")
	tmpl, err = tmpl.ParseFile(tmpFile.Name())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	buf := new(bytes.Buffer)
	if err := tmpl.ApplyYaml([]byte("port: 3000\nhost: http://localhost"), buf); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	ref := "host: http://localhost\nport: 3000\nname: it's a test"
	if buf.String() != ref {
		t.Errorf("expected: %s, got: %s", ref, buf.String())
	}
}

func TestApplyYaml(t *testing.T) {
	tmpl := New("test")
	tmpl, err := tmpl.Parse("host: {{.host}}\nport: {{.port}}\nname: it's a test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	buf := new(bytes.Buffer)
	if err := tmpl.ApplyYaml([]byte("port: 3000\nhost: http://localhost"), buf); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	ref := "host: http://localhost\nport: 3000\nname: it's a test"
	if buf.String() != ref {
		t.Errorf("expected: %s, got: %s", ref, buf.String())
	}
}

func TestApplyYamlNil(t *testing.T) {
	tmpl := New("test")
	tmpl, err := tmpl.Parse("host: http://localhost\nport: 3000\nname: it's a test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	buf := new(bytes.Buffer)
	if err := tmpl.ApplyYaml(nil, buf); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	ref := "host: http://localhost\nport: 3000\nname: it's a test"
	if buf.String() != ref {
		t.Errorf("expected: %s, got: %s", ref, buf.String())
	}
}

func TestApply(t *testing.T) {
	tmpl := New("test")
	tmpl, err := tmpl.Parse("host: {{.host}}\nport: {{.port}}\nname: it's a test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	rdBuf := bytes.NewBufferString("port: 3000\nhost: http://localhost")
	buf := new(bytes.Buffer)
	if err := tmpl.Apply(rdBuf, buf); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	ref := "host: http://localhost\nport: 3000\nname: it's a test"
	if buf.String() != ref {
		t.Errorf("expected: %s, got: %s", ref, buf.String())
	}
}

func TestFunc(t *testing.T) {
	content := []byte("I love Golang")
	fd, err := ioutil.TempFile("", "input")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	defer os.Remove(fd.Name())
	if _, err := fd.Write(content); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := fd.Close(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	tmpl := New("test")
	tmpl, err = tmpl.Parse("name: {{base64 (readfile .file)}}")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	rdBuf := bytes.NewBufferString(fmt.Sprintf("file: %s", fd.Name()))
	buf := new(bytes.Buffer)
	if err := tmpl.Apply(rdBuf, buf); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	ref := fmt.Sprintf("name: %s", base64.StdEncoding.EncodeToString(content))
	if buf.String() != ref {
		t.Errorf("expected: %s, got: %s", ref, buf.String())
	}
}
