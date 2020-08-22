#### Function w/ Multiple Return
- [Doc](https://golang.org/doc/effective_go.html#multiple-returns)
- In the last two functions, we only returned 1 value
- Sometimes we want to return 2 values
    - At the moment I'm writing this I'm kinda aware of the "errors"
    - So I'll probably write almost the same thing later, but like this:
    - _"Most of the times you should return +2 values, 1 of them being an __'Error'__ type"_
    - For now, you don't have to care that much about this topic
    - Later we'll be seeing it anyway
- The principal difference will be at the:
    - _returnType_
    - _return variable_
-   ```go
        package main
        
        import "fmt"
        
        func multipleReturn(num int, words ...string) (int, string)  {
            /* The first difference is at returnType
             * Before we've seen only 'string', or 'int'
             * In this case, we can add n-types inside the ()
            */
            /* This function will be:
             * Factorial
             * String concatenate
            */
            fac := 1
            for i := 1; i <= num; i++ {
                fac *= i
            }
            phrase := ""
            for _, word := range words {
                phrase += word + " "
            }
        
            /* Here is another difference
             * like the returnType, we'll be return
             * the values separated by a comma ","
             * Remember:
             * FOLLOWING THE SAME ORDER OF returnType
            */
            return fac, phrase
        }
        
        func main()  {
        //  (int, string) remember the order?
            factorial, sentence := multipleReturn(8, "Hello!", "This", "is", "a", "multiple", "return", "function")
        
            fmt.Println(sentence)
            fmt.Println("Factorial:", factorial)
        }
    ```