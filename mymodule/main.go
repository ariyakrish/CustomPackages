// main.go
package main

import (
    "fmt"
    "mymodule/calculator" // import path needs to match the module path and subdirectory
)

func main() {
    fmt.Println(calculator.Add(5, 7))        // Outputs: 12
    fmt.Println(calculator.Multiply(5, 7))   // Outputs: 35
}
