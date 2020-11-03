package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type HTMLData struct {
	Title string
	H2    string
	Ul    []Li
}

// Struct to generate a list
type Li struct {
	InnerText string
	TextColor string
	BGColor   string
}

// Literally writing pure HTML, nothing advanced
func getTemplate() string {
	/* Now we'll be injecting data into
	 * our html template
	 * each {{ . }} refers to an item
	 * so if there's a {{ .Foo }}
	 * this will be referring to a data named _Foo_
	 */
	t := `
	<html>
		<head>
			<title>{{ .Title }}</title>
		</head>
		<body>
			<h2>{{ .H2 }}</h2>
			{{ if .Ul }}			
			<ul>
				{{ range .Ul }}
					<li style="color: {{ .TextColor }}; background-color: {{ .BGColor }}">
						{{ .InnerText }}
					</li>	
				{{ end }}
			</ul>
			{{ end }}
		</body>
	</html>
		 `
	// Return html string
	return t
}

// Building HTML template
func getHtml() string {

	/* This will store the template buffer
	 * so it can be converted to string later
	 */
	var tpl bytes.Buffer

	// Defining our data
	data := HTMLData{
		Title: "Go Templatesl",
		H2:    "Successfully inject text",
		Ul: []Li{
			{
				InnerText: "Black",
				TextColor: "#FFFFFF",
				BGColor:   "#000000",
			},
			{
				InnerText: "Blue",
				TextColor: "#000000",
				BGColor:   "#0000FF",
			},
			{
				InnerText: "Red",
				TextColor: "#FFFFFF",
				BGColor:   "#FF0000",
			},
		},
	}

	// Parsing html string
	t, err := template.New("hello").Parse(getTemplate())
	checkErr(err)

	/* Writing Template into bytes.Buffer
	 * and injecting the data struct
	 * into template tags {{ .Foo }}
	 */
	err = t.Execute(&tpl, data)

	checkErr(err)

	// Converting bytes.Buffer to String
	return tpl.String()
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
