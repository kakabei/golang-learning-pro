package main 

import (

        "fmt"
        "os"
        "path/filepath"
)

func main(){

        sep := os.PathSeparator 
        fmt.Println(string(sep))

        fmt.Println(filepath.ToSlash(`c:\python\python.exe`))
}