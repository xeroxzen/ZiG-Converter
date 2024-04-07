// ZiG Currency Converter test suite


package main

import (
"testing"
)


func TestPerformConversion(t *testing.T) {
    testCases := []struct {
        name          string
        amount        float64
        fromCurrency  string
        toCurrency    string
        expected      float64
        shouldError   bool
    }{
        {"USD to ZWL", 100, "USD", "ZWL", 3390300, false},
        {"USD to ZiG", 100, "USD", "ZiG", 1356.16, false},
        {"ZWL to ZiG", 2500, "ZWL", "ZiG", 1, false},
        {"ZiG to USD", 1356.16, "ZiG", "USD", 100, false},
        {"Unsupported Currency", 100, "EUR", "ZWL", 0, true}, 
        {"Invalid Conversion", 100, "USD", "GBP", 0, true}, 
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := performConversion(tc.amount, tc.fromCurrency, tc.toCurrency)

            if tc.shouldError && err == nil {
                t.Error("Expected an error, but got none")
            }

            if !tc.shouldError && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }

            if !tc.shouldError && result != tc.expected {
                t.Errorf("Expected %.2f, got %.2f", tc.expected, result)
            }
        })
    }
}
