#### Custom Types
- In Go there are types that the developer creates
- This is something that is in the future, for now its more like an abstract idea
```go
// type myType subType
type myType int
var b myType
func main() {
    a := 10
    fmt.Printf("%T", a)
    fmt.Printf("\n%T", b)
    /**
    * a = b this is wrong
    * because 'a' is type 'int'
    * and b is type 'myType'
    **/
}
```
- This program should output:
```
int
main.myType
```
- Even with the same values "*a(int)*" != "*b(myType)*", because of the var type
