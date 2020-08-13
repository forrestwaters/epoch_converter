package main

import (
    "fmt"
    "time"
    "os"
    "strconv"
)

func main() {

    expectedUsage := "Usage: epoch 12345678"

    if len(os.Args) != 2 {
        fmt.Println(expectedUsage)
        os.Exit(1)
    }

    i, err := strconv.ParseInt(os.Args[1], 10, 64)

    if err != nil {
        fmt.Println(err)
        fmt.Println(expectedUsage)
        os.Exit(1)
    }

    time := time.Unix(i, 0)

    fmt.Println(time)

    os.Exit(0)

}
