package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

// Literally writing pure HTML, nothing advanced
func getTemplate() string {
	t := `
	<html>
		<head>
			<title>Teste</title>
		</head>
		<body>
			<h2>Hello, world!</h2>
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

	// Parsing html string
	t, err := template.New("hello").Parse(getTemplate())
	checkErr(err)

	// Writing Template into bytes.Buffer
	err = t.Execute(&tpl, nil)

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
