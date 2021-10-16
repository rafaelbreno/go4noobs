#### Println()
- I think it's a good habit read the [Official Documentation](https://golang.org/doc/) of any lang/framework/tool that you're using
- This case is a good example of that
- Here is the [Println() Func Doc](https://golang.org/pkg/fmt/#Println), reading it's possible to do this:
```go
func main() {
  //Here is simple function
  bytes, err := fmt.Println("Hello World,", "I'm Rafael Breno!")
  fmt.Println("The string has", bytes, "bytes. Error:", err)
}
```
- How do I know what's returning, and in what order? Following the doc:
    - "*It returns the number of bytes written and any write error encountered.*"
