package template

import (
	"html/template"
	"io"
)

type User struct {
	FirstName string
	LastName  string
	Prefix    string
}

const tmp = `Hello {{.Prefix}} {{.FirstName}} {{.LastName}}!`

func renderTemplate(out io.Writer, data User) error {
	t := template.Must(template.New("letter").Parse(tmp))
	return t.Execute(out, data)
}

func renderTemplateByMap(out io.Writer, data map[string]string) error {
	t := template.Must(template.New("letter").Parse(tmp))
	return t.Execute(out, data)
}
