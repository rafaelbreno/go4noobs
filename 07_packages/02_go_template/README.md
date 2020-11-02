### Go Template
- There're 2 packages of _templates_
    - [text/template](https://golang.org/pkg/text/template/)
    - [html/template](https://golang.org/pkg/html/template/)
- In this case I'll be sticking with the _html_ package
- To show the built page the package _net/http_ should be imported
- The base of this example will be
-   ``go
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

    ``
    - _checkErr:_ Simple helper to panic errors
    - _test:_ Func that will handle http access and draw the HTML
- Now with the base we can start with the _html/template_
#### _Level 1_
- For this level the best way to start would be the famous _Hello, World!_
-   ```go
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
    ``` 
- This will generate a simple webpage at _localhost:3000_ 
#### Level 2
- In this level we'll see some data
-   ``go

    ``
