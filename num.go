package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {

// based on: http://piratepad.net/XJnL7LdAUx

    var inputNumber int
    var count int
    var ageInput int
    var input string
    var fileName string

    for checkLogin() == false {
    }

    fmt.Println("Enter your age. Be honest. Don't fuck with me.")
    fmt.Scanln(&ageInput)

    if ageInput < 18 {
       fmt.Println("You shalt not pass.")
       return
    }

    fmt.Println("Enter the number you want to print.")
    fmt.Println("Your input will be incrementally printed starting with 1.")
    fmt.Println("Enter desired number.")
    fmt.Scanln(&inputNumber)

    for count != inputNumber {
        count++
// no longer print to stdout, let's write to /tmp/nondiv.txt instead!
        fmt.Println(count)
        if count%5 == 0 {
            writeFile(count)
        }
    }

    for input == "" || fileName == "" {
        fmt.Println("Tell me all of your secrets.")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        input = scanner.Text()
	    if err := scanner.Err(); err != nil {
	    	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	    }
        fmt.Println("Alright, which Filename should we save your file to?")
        fmt.Scanln(&fileName)
    }
    writeFileWithName(input, fileName)
}

func checkLogin() bool {
    var userAccountInput string
    var userPasswordInput string
    var userAccountStatic string = "fucknut"
    var userPasswordStatic string = "nutfuck"

    fmt.Println("Enter your username: ")
    fmt.Scanln(&userAccountInput)

    fmt.Println("Enter your password: ")
    fmt.Scanln(&userPasswordInput)

    if userAccountInput == userAccountStatic && userPasswordInput == userPasswordStatic {
        fmt.Println("Access granted.")
        return true
    } else {
        fmt.Println("Nice try. Now die in a fire, would you?")
        return false
    }
}

func writeFile(num int) {
    fmt.Println("Writing non-divideable numbers to file: " +strconv.Itoa(num))
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

func writeFileWithName(input string, fileName string) {
    fmt.Println("Writing input to file: ")
    f, err := os.OpenFile("/tmp/"+fileName, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Error #1: Couldn't write to file.")
        fmt.Println(err)
    }
    n, err := f.WriteString(input)
    fmt.Println("Writing to /tmp/"+fileName)
    fmt.Println("Done.")
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
