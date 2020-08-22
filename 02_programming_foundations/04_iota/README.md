#### Iota
- [Doc](https://golang.org/ref/spec#Iota)
- By _first sight_ looks like a counter inside const's scope
    - ```
        const (
            a = iota // iota == 0
            b = iota // iota == 1
            c = iota // iota == 2
            d = iota // iota == 3
        )
      ```