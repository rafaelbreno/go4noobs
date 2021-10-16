#### Slice of Slice
- [Doc](https://golang.org/doc/effective_go.html#slices)
- To deal with a Slice of a Slice, it's need to use the following formats:
    - ``slice = slice[:]``
    - ``slice = slice[a:]``
    - ``slice = slice[:b]``
    - ``slice = slice[a:b]``
- The most _"educational"_ way to explain what the _"a"_ and _"b"_ is, is with numerical sets/intervals
    - _"a"_ is left-closed
    - _"b"_ is right-open
    - For example, we have a numerical set/interval R = [1, 10]
        - R = 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 
    - If I asked you to select the interval from 3 to 8, you would do something like this
        - _[3, 8]_
        - _right-close_ and _left-closed_
    - _But_ if I asked you to select the interval from 3 to 8, but _right-open_ and _left-closed_, you would need to do this:
        - _[3, 9[_
        - The __"9["__ won't include _"9"_
```go
//                 0  1  2  3  4  5  6
var slice = []string{"a", "b", "c", "d", "e", "f", "g"}

func main()  {
    fmt.Println("-----[:]-----")
    slice1 := slice[:]
    for key, value := range slice1 {
        fmt.Println(key, ":", value)
    }
    fmt.Println("\n-----[3:]-----")
    slice2 := slice[3:]
    for key, value := range slice2 {
        fmt.Println(key, ":", value)
    }
    fmt.Println("\n-----[:3]-----")
    slice3 := slice[:3]
    for key, value := range slice3 {
        fmt.Println(key, ":", value)
    }
    fmt.Println("\n-----[2:5]-----")
    slice4 := slice[2:5]
    for key, value := range slice4 {
        fmt.Println(key, ":", value)
    }

    fmt.Println("\n-----[:2] U [4:]-----")
    /*
     * Again with math notations
     * remember "U"?
     * That means "union"
     * In this case [:2] U [4:]
     * Will be uniting the terms from 0 to 1, and from 4 until the last one
     * To unite them both we'll be using append with the "..." operator AGAIN!
    */
    union := append(slice[:2], slice[4:]...)
    fmt.Println(union)
}
```
