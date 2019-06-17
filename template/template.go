package template

import (
	"io"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Template is the parsed template structure.
type Template struct {
	tmpl *template.Template
}

// New parses input and creates a Template instance with given name.
func New(name string, in string) (*Template, error) {
	tmpl, err := template.New(name).Parse(in)
	if err != nil {
		return nil, err
	}
	return &Template{
		tmpl: tmpl,
	}, nil
}

// ApplyYaml applies yaml input to the template and write its output to writer.
func (t *Template) ApplyYaml(in []byte, wr io.Writer) error {
	arg := make(map[string]interface{})
	if err := yaml.Unmarshal(in, &arg); err != nil {
		return err
	}
	return t.tmpl.Execute(wr, arg)
}
