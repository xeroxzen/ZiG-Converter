**README.md**

## ZiG Currency Converter

A Go-based currency converter for Zimbabwe's new currency, ZiG, as well as the old currency, ZWL, and USD.

**Features**

- Converts between ZiG, ZWL, and USD.
- Utilizes up-to-date exchange rates (note how you'll handle keeping these rates updated).
- Handles invalid conversions and unsupported currencies.

**Getting Started**

**Prerequisites**

- Go installed on your system ([https://go.dev/doc/install](https://go.dev/doc/install))

**Installation**

1. Clone this repository:
   ```bash
   git clone git@github.com:xeroxzen/ZiG-Converter.git
   ```
2. Navigate to the project directory:
   ```bash
   cd ZiG-Converter
   ```

**Usage**

1. Build the executable:
   ```bash
   go build
   ```
2. Run the converter:
   ```bash
   ./converter.go
   ```
   You'll be prompted to enter the amount, original currency, and target currency Mind your typing on the currencies.

**Example**

```
Enter amount: 100
Enter original currency (USD, ZWL, ZiG): USD
Enter target currency (USD, ZWL, ZiG): ZiG
100 USD is equivalent to 1356.16 ZiG
```

**Updating Exchange Rates**

Currently, exchange rates are hardcoded. To update them:

1. Modify the `USD_TO_ZWL`, `USD_TO_ZIG`, and `ZWL_TO_ZIG` constants in `converter.go`.
2. Rebuild the executable.

**Future Feature:**

- Ability to fetch exchange rates from an online API.

**Contributing**

Contributions are welcome! Feel free to open issues or submit pull requests.

**Disclaimer**

This currency converter is intended for informational and educational purposes.

**License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**Acknowledgments**

The Newsday article ["Zimbabwe introduces new currency"](https://www.newsday.co.zw/opinion-analysis/article/200025272/here-is-how-to-convert-your-zimdollar-balances-to-zimbabwe-gold-zig) was used as a reference for the exchange rates.

**Author**

- [xeroxzen](https://github.com/xeroxzen)(
  Andile Jaden Mbele
  )
