package main

import (
    "fmt"
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
            fmt.Println(count)
        }
    }
}
