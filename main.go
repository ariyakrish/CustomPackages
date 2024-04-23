// mathutil.go

package CustomPackages

import  ( "os"
"log"
"fmt"

)

// Add takes two integers and returns their sum.
func Add(a, b int) int {
    return a + b
}

// Add takes two integers and returns their sum.
func CreateFile()  {
    _, err := os.Create("package.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File created:", "package.txt")

}
