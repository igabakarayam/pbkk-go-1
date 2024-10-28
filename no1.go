package main

import (
    "fmt"
)

func main() {
    var input int
    fmt.Print("Masukkan sebuah angka: ")
    fmt.Scan(&input)

    if input == 42 {
        fmt.Println("Hello World")
    } else {
        fmt.Println("Input Anda adalah:", input)
    }
}
