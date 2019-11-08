package yamltmpl

import (
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Template is the parsed template structure.
type Template struct {
	tmpl *template.Template
}

// New creates a Template instance.
func New(name string) *Template {
	return &Template{
		tmpl: template.New(name),
	}
}

// Parse the template input.
func (t *Template) Parse(in string) (*Template, error) {
	tmpl, err := t.Funcs(defaultFunc).tmpl.Parse(in)
	if err != nil {
		return nil, err
	}
	t.tmpl = tmpl
	return t, nil
}

// ParseFile parses the file as the template input.
func (t *Template) ParseFile(file string) (*Template, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	cfg, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return t.Parse(string(cfg))
}

// Funcs adds the function to the template's function map.
func (t *Template) Funcs(funcs template.FuncMap) *Template {
	t.tmpl.Funcs(funcs)
	return t
}

// ApplyYaml applies yaml input to the template and write its output to writer.
func (t *Template) ApplyYaml(in []byte, wr io.Writer, overrides ...map[string]interface{}) error {
	arg := make(map[string]interface{})
	if in != nil {
		if err := yaml.Unmarshal(in, &arg); err != nil {
			return err
		}
	}
	applyOverrides(arg, overrides...)
	return t.tmpl.Execute(wr, arg)
}

// Apply applies yaml input from Reader to the template and write its output to writer.
func (t *Template) Apply(rd io.Reader, wr io.Writer, overrides ...map[string]interface{}) error {
	arg := make(map[string]interface{})
	dec := yaml.NewDecoder(rd)
	if err := dec.Decode(&arg); err != nil {
		return err
	}
	applyOverrides(arg, overrides...)
	return t.tmpl.Execute(wr, arg)
}
