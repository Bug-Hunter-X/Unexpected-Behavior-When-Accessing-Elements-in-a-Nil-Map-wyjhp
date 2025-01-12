```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, but print 0
    fmt.Println(m["b"]) // This will not panic, but print 0
    m["a"] = 100
    fmt.Println(m["a"]) // This will print 100
}
```