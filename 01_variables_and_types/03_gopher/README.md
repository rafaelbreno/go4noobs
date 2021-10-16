#### Gopher
- Gopher is a short operator Its symbol is *":="* 

##### What it does
- [Official Doc - Short Variable Declarations](https://golang.org/ref/spec#Short_variable_declarations)
- The operator *:=* infer a type and assing a value to a variable
```golang
a := 10 // a of type int and value 10
b := "foo" // b of type string and value "foo"
```
- Different of *"="*, that is a *assignment operator*

##### Rules
- You can see more of it in these links:
    - [Unofficial Doc - Stackoverflow Question](https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go/45654233#45654233)
1. You can't use *":="*  out of *"funcs"* . It's because, out of any func, a statement should start with a keyword.
```go
// no keywords below, illegal.
illegal := 42

// `var` keyword makes this statement legal.
var legal = 42

func foo() {
  alsoLegal := 42
  // reason: it's in a func scope.
}
```
2. You can't use them twice (in the same scope):
```go
legal := 42
legal := 42 // <-- error
```
3. You can use them for multi-variable declarations and assignments:
```go
foo, bar   := 42, 314
jazz, bazz := 22, 7
```
4. You can use them twice in **"multi-variable"** declarations, if one of the variables is new:
```go
foo, bar  := someFunc()
foo, jazz := someFunc()  // <-- jazz is new
baz, foo  := someFunc()  // <-- baz is new
```
5. You can use the short declaration to declare a variable in a newer scope even that variable is already declared with the same name before:
```go
var foo int = 34

func some() {
  // because foo here is scoped to some func
  foo := 42  // <-- legal
  foo = 314  // <-- legal
}
```
6. You can declare the same name in short statement blocks like: **if**, **for**, **switch**:
```go
foo := 42
if foo := someFunc(); foo == 314 {
  // foo is scoped to 314 here
  // ...
}
// foo is still 42 here
```
- With it said, we can now use the function [Printf]() and the [Printing "verbs"](https://golang.org/pkg/fmt/#hdr-Printing)
- To help understand that *":="* guess what type of value it's assigned to it
```go
func main(){
  numInt := 20
  numFloat := 10.2
  textString := "Hello, my name is Rafael"
  textChar := 'R'
  fmt.Printf("Var: %b , Type: %T\n", numInt, numInt)
  fmt.Printf("Var: %b , Type: %T\n", numFloat, numFloat)
  fmt.Printf("Var: %b , Type: %T\n", textString, textString)
  fmt.Printf("Var: %b , Type: %T\n", textChar, textChar)
}
```
