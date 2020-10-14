#### Recursive Functions
- A recursive function is a function that finite re-uses itself
- ```go
    package main
    
    import "fmt"
    
    func main()  {
    	fmt.Println(fac(5)) // 120
    }
    
    // Factorial
    func fac(n int) int {
    	// The function "breaks" when n == 1
    	if n == 1 {
    		return 1
    	}
    	return fac(n - 1) * n
    }
  ```