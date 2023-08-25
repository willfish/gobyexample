package main

import (
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name  string
	Email string
}

func main() {
	tmpl := raw(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>{{.Name}}</title>
			</head>
		<body>
			<h1>{{.Name}}</h1>
			<p>Contact: {{.Email}}</p>
		</body>
		</html>
`)

	p := Person{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	t := template.Must(template.New("html").Parse(tmpl))
	t.Execute(os.Stdout, p)
}

func raw(s string) string {
	return strings.TrimLeft(s, "\n\t")
}
