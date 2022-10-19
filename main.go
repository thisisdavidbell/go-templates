package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name           string
	Age            int
	FavouriteFoods []string
}

// Following example and doc:
//   - https://pkg.go.dev/text/template

func main() {

	person1 := Person{
		Name: "David",
		Age:  1001,
	}
	hello_world_template := "Hello {{.Name}}. You are {{.Age}} years old.\n"

	personTemplate, err := template.New("tmpl").Parse(hello_world_template)
	if err != nil {
		panic(err)
	}

	// follow example straight to stdout
	err = personTemplate.Execute(os.Stdout, person1)
	if err != nil {
		panic(err)
	}

	// if our goal is to output to a string, not directly an io.Writer (like os.Stdout), we can use 2 approaches which implement io.Writer:
	//   - strings.Builder - preferred: https://dev.to/shiraazm/efficiently-concatenating-strings-in-go-6do
	//   - bytes.buffer
	fmt.Println("\n========== To String example ==========\n")

	// bytes.buffer
	var b bytes.Buffer
	err = personTemplate.Execute(&b, person1) // pointer needed
	if err != nil {
		panic(err)
	}
	fmt.Printf("string from template.Execute to bytes.buffer: %v\n", b.String())

	//strings.Builder
	var sb strings.Builder
	err = personTemplate.Execute(&sb, person1) // pointer needed
	if err != nil {
		panic(err)
	}
	fmt.Printf("string from template.Execute to strings.Builder: %v\n", sb.String())

	fmt.Println("\n========== White Space example ==========\n")

	white_space_template := "Hello    {{- .Name -}}    . You are     {{- .Age -}}     years old.\n"

	personTemplate, err = template.New("tmpl").Parse(white_space_template)
	if err != nil {
		panic(err)
	}

	err = personTemplate.Execute(os.Stdout, person1)
	if err != nil {
		panic(err)
	}

	// Try new actions from:
	/* Things to cover before finishing:
	- [x] if
	- [ ] gt
	- len
	- [x] range

	Other things to try:
	- [ ] applying functions
	*/

	fmt.Println("\n========== Comments example ==========\n")

	comment_template := "White space left here    {{/* IM THE COMMENT */}}    but no sign of comment.\n" // note: no ws between {{ and /*
	personTemplate, err = template.New("tmpl").Parse(comment_template)
	if err != nil {
		panic(err)
	}
	err = personTemplate.Execute(os.Stdout, person1)
	if err != nil {
		panic(err)
	}

	comment_template = "White space gone    {{- /* IM THE COMMENT */ -}}    but no sign of comment.\n"
	personTemplate, err = template.New("tmpl").Parse(comment_template)
	if err != nil {
		panic(err)
	}
	err = personTemplate.Execute(os.Stdout, person1)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n========== If example ==========\n")

	if_template := "Show name: {{.Name}}\n"
	personTemplate, err = template.New("tmpl").Parse(if_template)
	if err != nil {
		panic(err)
	}
	err = personTemplate.Execute(os.Stdout, person1)
	if err != nil {
		panic(err)
	}

	if_template = "Is Name Empty: {{if .Name}} No {{end}}\n"
	personTemplate, err = template.New("tmpl").Parse(if_template)
	if err != nil {
		panic(err)
	}
	err = personTemplate.Execute(os.Stdout, person1)
	if err != nil {
		panic(err)
	}

	emptyPerson := Person{}

	if_template = "Is Name Empty: {{if .Name}} No {{else}} Yes {{end}}\n"
	personTemplate, err = template.New("tmpl").Parse(if_template)
	if err != nil {
		panic(err)
	}
	err = personTemplate.Execute(os.Stdout, emptyPerson)
	if err != nil {
		panic(err)
	}

	if_template = "Is Name Empty: {{if .Name}} No, its {{.Name}} {{end}}\n"
	personTemplate, err = template.New("tmpl").Parse(if_template)
	if err != nil {
		panic(err)
	}
	err = personTemplate.Execute(os.Stdout, person1)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n========== Range example ==========\n")

	// people := []Person{
	// 	{
	// 		Name: "name1",
	// 		Age:  1,
	// 	},
	// 	{
	// 		Name: "name2",
	// 		Age:  2,
	// 	},
	// 	{
	// 		Name: "name3",
	// 		Age:  3,
	// 	},
	// }
	foodPerson := Person{
		Name:           "FoodLover",
		Age:            21,
		FavouriteFoods: []string{"curry", "smoothies", "fruit", "ice cream"},
	}
	range_template := "Range example. This persons fav foods are:\n {{range .FavouriteFoods}} - {{.}} \n {{end}}\n"
	personTemplate, err = template.New("rangeTmpl").Parse(range_template)
	if err != nil {
		panic(err)
	}
	err = personTemplate.Execute(os.Stdout, foodPerson)
	if err != nil {
		panic(err)
	}
}
