package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type HTMLData struct {
	Title string
	H1    string
}

var tpl bytes.Buffer

/*
 *
**/
func getRawTemplate() string {
	return `
		<html>
			<head>
				<title>{{ define "Title" }}{{ .Title }}{{end}}</title>
			</head>
			<body>
				<h1>{{ template "Header1" }}</h1>
			</body>
		</html>
	`
}

func buildTemplate() *template.Template {
	t, err := template.New("main").Parse(getRawTemplate())
	checkErr(err)
	return t
}

func getData() HTMLData {
	return HTMLData{
		Title: "Lorem Ipsum",
		H1:    "Go Templates Example",
	}
}

func yield(t *template.Template, section string, value string) *template.Template {
	/* With this function we can
	 *insert values inside templates
	 */
	t.New(section).Parse(value)
	return t
}

func getBuiltTemplate(t *template.Template) string {

	err := t.Execute(&tpl, getData())
	checkErr(err)

	return tpl.String()
}

func getHtml() string {
	t := buildTemplate()

	t = yield(t, "Header1", "Allo aloo")

	return getBuiltTemplate(t)
}

/*---------- App base ---------*/

// Helper to panic errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Func to handle the route
func test(w http.ResponseWriter, req *http.Request) {
	// Return HTML
	fmt.Fprintf(w, getHtml())
}

// Main func
func main() {
	http.HandleFunc("/", test)

	http.ListenAndServe(":3000", nil)
}
