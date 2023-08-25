package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")

	t1, err := t1.Parse("hello {{.}}!")

	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "world")
	t1.Execute(os.Stdout, 123)
	t1.Execute(os.Stdout, true)
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "hello {{.Name}}!\n")

	t2.Execute(
		os.Stdout,
		struct{ Name string }{"world"},
	)

	t2.Execute(
		os.Stdout,
		map[string]string{"Name": "world"},
	)

	t3 := Create(
		"t3",
		"{{if . -}} yes {{else -}} no {{end}}\n",
	)

	t3.Execute(os.Stdout, true)
	t3.Execute(os.Stdout, false)

	t4 := Create(
		"t4",
		"Range:{{range .}} {{.}} {{end}}\n",
	)

	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
}
