#### Procedural Programming
- Procedural programming is a programming paradigm
    - The simplest way to explain it is: the program will be executed line by line
```go 
func main() {
  fmt.printf(var01) // flag 01
  var01 := "hello coder!!" // flag 02
  fmt.printf(var01) // flag 03
}
```
- If you run the code above will get an error, why?
- Because you printed the variable before declaring it
    - At the _Flag 01_ the program don't "know" the existence of the variable __var01__ _yet_
        - Here the program will brake, but just for study purposes we will continue the logical processes
    - Then at the _Flag 02_ where finally the __var01__ were declared the program will know the existence of it
    - Finally, at _Flag 03_ the program will be able to print __var01's__ value
- __Disclaimer!!!__ This is the basic explanation about procedural paradigm
- __I truly__ recommend you reserve a tiny bit of your day to give it a try.
- Just to search about some "rules" and "good behavior" when programming in _Go_
