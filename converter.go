/*
Program: ZiG Currency Converter
Author: Andile Jaden Mbele
Date: 7 April 2024
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

const (
    USD_TO_ZWL = 33903.0
    USD_TO_ZIG = 13.5616
    ZWL_TO_ZIG = USD_TO_ZWL / USD_TO_ZIG
)

func getUserInput() (float64, string, string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter amount: ")
    amountStr, _ := reader.ReadString('\n')
    amount, err := strconv.ParseFloat(amountStr[:len(amountStr)-1], 64)
    if err != nil {
        fmt.Println("Invalid amount")
        os.Exit(1)
    }

    fmt.Print("Enter original currency (USD, ZWL, ZiG): ")
    fromCurrency, _ := reader.ReadString('\n')
    fromCurrency = fromCurrency[:len(fromCurrency)-1]

    fmt.Print("Enter target currency (USD, ZWL, ZiG): ")
    toCurrency, _ := reader.ReadString('\n')
    toCurrency = toCurrency[:len(toCurrency)-1]

    return amount, fromCurrency, toCurrency
}

func performConversion(amount float64, fromCurrency, toCurrency string) (float64, error) {
    // Conversion logic based on currency pairs
    switch fromCurrency {
    case "USD":
        switch toCurrency {
        case "ZWL":
            return amount * USD_TO_ZWL, nil
        case "ZiG":
            return amount * USD_TO_ZIG, nil
        default:
            fmt.Println("Unsupported conversion")
            return 0, fmt.Errorf("unsupported conversion from %s to %s", fromCurrency, toCurrency)
        }
    case "ZWL":
        switch toCurrency {
        case "USD":
            return amount / USD_TO_ZWL, nil
        case "ZiG":
            return amount / ZWL_TO_ZIG, nil
        default:
            fmt.Println("Unsupported conversion")
            return 0, fmt.Errorf("unsupported conversion from %s to %s",  fromCurrency, toCurrency)
    }
    case "ZiG":
        switch toCurrency {
        case "USD":
            return amount / USD_TO_ZIG, nil        
        case "ZWL":
            return amount * ZWL_TO_ZIG, nil
        default:
            fmt.Println("Unsupported conversion")
            return 0, fmt.Errorf("unsupported conversion from %s to %s", fromCurrency, toCurrency)
    }
    default:
        fmt.Println("Unsupported conversion")
        return 0, fmt.Errorf("unsupported conversion from %s to %s", fromCurrency, toCurrency)
    }
}

func main() {
    amount, fromCurrency, toCurrency := getUserInput()
    result, err := performConversion(amount, fromCurrency, toCurrency)
    if err != nil {
        os.Exit(1)
    }
    fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, fromCurrency, result, toCurrency)
}
