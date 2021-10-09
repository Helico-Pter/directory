package main

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {

    files, err := filepath.Glob("/home/*/**/*.txt")


    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {

        fmt.Println(file)
    }
}
