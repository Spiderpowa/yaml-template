package yamltmpl

import (
	"bytes"
	"testing"
)

func TestApplyYaml(t *testing.T) {
	tmpl, err := New("test", "host: {{.host}}\nport: {{.port}}\nname: it's a test")
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

func TestApply(t *testing.T) {
	tmpl, err := New("test", "host: {{.host}}\nport: {{.port}}\nname: it's a test")
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
