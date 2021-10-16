#### String
- [Doc](https://golang.org/ref/spec#String_literals)
- String is an array os _characters_
```go
package main
import "fmt"
func main()  {
    s := "Hello, World!"
    sb := []byte(s)

    for _, v := range sb {
        fmt.Printf("%v - %T <-> %#U - %#x\n", v, v, v, v)
    }

    // Char per char
    for _, z := range s {
        fmt.Printf("%v - %T <-> %#U - %#x\n", z, z, z, z)
    }

    // Byte per Byte
    for i := 0; i < len(s); i++ {
        fmt.Printf("%v - %T <-> %#U - %#x\n", s[i], s[i], s[i], s[i])
    }
}
```
- The _%v_ is already known, but what's the _#_ in the middle of _%#U_ and _%#x_
    - __#__ is known as _Flag_, it's 1 in 5
- _[Flags:](https://golang.org/pkg/fmt/)_
    - _+ :_ Always print a sign for numeric values
    - _- :_ Pads with space rather on left than right
    - _# :_ Alternate the format
    - _' '(space):_ Leave a space for elided sign on numbers _(% b)_
    - _0 :_ Pad with leading zeros
- [More about string](https://blog.golang.org/strings)
