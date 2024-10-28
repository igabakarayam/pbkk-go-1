package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Fungsi untuk membalikkan sebuah kata
func reverseWord(word string) string {
    runes := []rune(word)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Masukkan sebuah kalimat (minimal 3 kata): ")
        scanner.Scan()
        input := scanner.Text()

        words := strings.Fields(input) // Menghitung jumlah kata
        if len(words) < 3 {
            fmt.Println("Input ditolak. Minimal 3 kata diperlukan.")
            continue
        }

        // Membalikkan setiap kata di dalam kalimat
        for i, word := range words {
            words[i] = reverseWord(word)
        }

        reversedSentence := strings.Join(words, " ")
        fmt.Println("Hasil reverse per kata:", reversedSentence)
        break
    }
}
