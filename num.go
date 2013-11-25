package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {

    fmt.Println("Enter the number you want to print.")
    fmt.Println("Your input will be incrementally printed starting with 1.")

    var inputNumber int
    var count int

    fmt.Println("Enter desired number.")
    fmt.Scanln(&inputNumber)

    for count != inputNumber {
        count++
        if count%5 == 0 {
// no longer print to stdout, let's write to /tmp/nondiv.txt instead!
//            fmt.Println(count)
        writeFile(count)
        }
    }
}

func writeFile(num int) {
    fmt.Println("Writing non-divideable numbers to file: ")
    f, err := os.OpenFile("/tmp/nondiv.txt", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Error #1: Couldn't write to file.")
        fmt.Println(err)
    }
    n, err := f.WriteString(strconv.Itoa(num) + "\n")
    if err != nil {
        fmt.Println("Error #2: Couldn't write fo file.")
        fmt.Println(n, err)
    }
    f.Close()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
