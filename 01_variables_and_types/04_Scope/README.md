#### Scope
- Scope is an area with limits defined by ``{}``
    - For Example``{ ...in scope... } ...out of scope`` 
- A good example of it is a "Global Variable", its **scope** is the entire code, can be used in every *func* of a package
- At the code below we can see that the variable **global** can be used 
- ```go
      var global = 10
      func main() {
        local := 20
        fmt.Print("Global Variable:", global)
        fmt.Print("\nLocal Variable:", local)
        Scope(local)
      }     
      func Scope(x int){
        //This is legal
        fmt.Print("\n\nGlobal Variable:", global)
      
        //This is legal
        fmt.Print("\nParameter Variable:", x)
      
        //This is illegal, will show 'undefined' error in terminal
        // fmt.Print("Main Variable:", local)
      }
  ```