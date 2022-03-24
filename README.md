# genset

A Generic Set implementation for Go.

```bash
go get github.com/markekraus/genset
```

```go
package main

import (
    "fmt"

    genset "github.com/markekraus/genset/pkg"
)

type mytype struct {
    a, b int
}

func main() {
    s := genset.New[int]()
    s.Add(3)
    s.Add(2)
    s.Add(90)
    s.Add(90)
    fmt.Printf("length: %v\n", s.Len())
    fmt.Printf("has 3: %v\n", s.Has(3))
    fmt.Printf("has 50: %v\n", s.Has(50))
    s.Remove(2)
    fmt.Printf("length: %v\n", s.Len())
    fmt.Printf("has 2: %v\n", s.Has(2))

    s.AddMulti(5, 6, 90)
    fmt.Printf("length: %v\n", s.Len())

    filtered := s.Filter(func(value int) bool {
        return value < 50
    })
    fmt.Printf("length: %v\n", filtered.Len())
    fmt.Printf("has 3: %v\n", filtered.Has(3))
    fmt.Printf("has 90: %v\n", filtered.Has(90))

    abort := make(chan struct{})
    for v := range s.Range(abort) {
        fmt.Printf("value: %v\n", v)
        if v == 5 {
            close(abort)
            break
        }
    }
}
```

Output:

```plaintext
length: 3
has 3: true
has 50: false
length: 2
has 2: false
length: 4
length: 3
has 3: true
has 90: false
value: 3
value: 5
```
