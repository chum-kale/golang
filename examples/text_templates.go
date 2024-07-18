package main

import (
	"html/template"
	"os"
)

func main() {
	//create template and parse it's body from a string
	//Templates are a mix of static text and “actions” enclosed in {{...}} that are used to dynamically insert content
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	//panic when parse fails
	//good for global templates
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	//helper function
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	//templates for maps
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	//if-else using templates
	//drimming whitespaces//conditional templates
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	//using range
	//Inside the range block {{.}} is set to the current item of the iteration.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
