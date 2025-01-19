```go
package main

import (

    "fmt"
)

func main() {
    var m map[string]int
    // Explicitly check if the map is initialized
    if m == nil {
        fmt.Println("Map is not initialized")
        return
    }

    value, ok := m["a"]
    if ok {
        fmt.Println("Value of a:", value)
    } else {
        fmt.Println("Key "a" not found")
    }
    
    // Alternatively, initialize the map
    m = make(map[string]int)
    fmt.Println(m["a"]) // Prints 0 as the map is initialized now
    m["a"] = 10
    fmt.Println(m["a"]) //Prints 10
}
```