package main

import (
    "fmt"
)

func main() {

    fmt.Println("Enter the number you want to print.")
    fmt.Println("Your input will be incrementally printed starting with 1.")

    var inputNumber int
    var countToInput int

    fmt.Println("Enter desired number.")
    fmt.Scanln(&inputNumber)

    for countToInput != inputNumber {
        countToInput++
        fmt.Println(countToInput)
    }
}
