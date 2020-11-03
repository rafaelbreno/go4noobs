### Go Template
- There're 2 packages of _templates_
    - [text/template](https://golang.org/pkg/text/template/)
    - [html/template](https://golang.org/pkg/html/template/)
- In this case I'll be sticking with the _html_ package
- To show the built page the package _net/http_ should be imported
- The base of this example will be
-   ```go
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

    ```
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
#### _Level 2_
- In this level we'll see some data injection
-   ```go
        type HTMLData struct {
        	Title string
        	H2    string
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
    ```

- Inside of our html template there is some _"tags"_
- `{{ .Title }}` will receive `data.Title` value
- `{{ .H2 }}` will receive `data.H2` value
#### _Level 3_
- Now with we can add more tags

-   ```go
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
    ```
#### _Level 4_
- Now we'll inject some html snippets inside the template
- Almost like [Laravel Blade](https://laravel.com/docs/master/blade) _@yield() / @sections()_ tags
-   ```go

    ```

