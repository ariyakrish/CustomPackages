// mathutil.go

package CustomPackages

import (
    "os"
)

// Add takes two integers and returns their sum.
func Add(a, b int) int {
    return a + b
}
func Sub(a, b int) int {
    return a - b
}

func CreateFile() string{
    _, err := os.create("Sample.txt")
    if err!= nil{
        return err
    }
    return "File Created Succesfully"
}