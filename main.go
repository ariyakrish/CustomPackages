// mathutil.go

package CustomPackages

import  ( "os"
"log"

)


// Add takes two integers and returns their sum.
func CreateFile()  {
    f, err := os.Create("package.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    fmt.Println("File created:", "package.txt")

}
