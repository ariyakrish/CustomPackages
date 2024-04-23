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
 
func CreateFile(FileName string) (*os.File, error) {
    // Check if the file exists
    _, err := os.Stat(FileName)
    if err == nil {
        return nil, os.ErrExist // File already exists
    }
    if !os.IsNotExist(err) {
        return nil, err // Unexpected error occurred while checking file existence
    }
 
    // File doesn't exist, proceed to create it
    file, err := os.Create(FileName)
    if err != nil {
        return nil, err
    }
 
    // Ensure the file is closed when the function returns
    defer file.Close()
 
    return file, nil
}