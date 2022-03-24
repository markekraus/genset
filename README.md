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
    fmt.Printf("items: %v, want: %v\n", s.Len(), 0)
    s.Add(3)
    s.Add(2)
    s.Add(90)
    s.Add(90)
    fmt.Printf("items: %v, want: %v\n", s.Len(), 3)
    fmt.Printf("has 3: \n", s.Has(3))
    fmt.Printf("has 50: \n", s.Has(50))
}
```

Output:

```plaintext
has 3: true
has 50: false
```
