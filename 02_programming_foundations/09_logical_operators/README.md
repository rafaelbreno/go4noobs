#### Conditionals
- [Doc](https://golang.org/ref/spec#Logical_operators)
- There are 3 logical operators
    - __&&__ (AND)
        - A _&&_ B = _"if A then B, else false"_
        - |    A|    B|A _&&_ B|
          |-----|-----|--------|
          |false|false|   false|
          |false|true |   false|
          |true |false|   false|
          |true |true |    true|
      
    - __||__ (OR)
        - A _||_ B = _"if A then true else B"_
        - |    A|    B|A _\\\\_ B|
          |-----|-----|--------|
          |false|false|   false|
          |false|true |    true|
          |true |false|    true|
          |true |true |    true|
    - __!__ (NOT)
        - _!_ A = _"not A"_
        - |A    |   !A|
          |-----|-----|
          |false| true|
          |true |false|
- __DISCLAIMER__
    - The function ``strconv.FormatBool`` receive a boolean as parameter and return it in string form
    - For example _strconv.FormatBool(true)_ will return _"true"_
-   ```
        package main
        
        import (
            "fmt"
            "strconv"
        )
        
        /*
        * In this program I'll be developing
        * a program that will build the same
        * tables shown at README.md file
        */
        func main()  {
            var values [2]bool
            values[0] = false
            values[1] = true
            /*
             * values = [false, true]
             */
            fmt.Println("---And Operator---")
            for _, valueA := range values {
                for _, valueB := range values {
                    fmt.Printf("%s && %s = %s\n",
                        strconv.FormatBool(valueA),
                        strconv.FormatBool(valueB),
                        strconv.FormatBool(valueA && valueB))
                }
            }
            fmt.Println("---Or Operator---")
            for _, valueA := range values {
                for _, valueB := range values {
                    fmt.Printf("%s || %s = %s\n",
                        strconv.FormatBool(valueA),
                        strconv.FormatBool(valueB),
                        strconv.FormatBool(valueA || valueB))
                }
            }
            fmt.Println("---Not Operator---")
            for _, valueA := range values {
                fmt.Printf("!%s = %s\n",
                    strconv.FormatBool(valueA),
                    strconv.FormatBool(!valueA))
            }
        }
    ```