package main // Define the package for this file

import ( // Import necessary packages
	"encoding/json" // For JSON encoding and decoding
	"fmt"           // For formatted I/O, similar to printf and scanf
	"log"           // For logging and handling errors
	"net/http"      // For making HTTP requests
	"strconv"       // For string conversion
	"strings"       // For string manipulation

	"github.com/spf13/cobra" // Cobra library for building CLI applications
)

const apiURL = "https://api.exchangerate-api.com/v4/latest/" // Constant for the API URL

type exchangeRateResponse struct { // Define a struct for the JSON response
	Rates map[string]float64 `json:"rates"` // Map to hold currency rates
}

func main() { // Main function, entry point of the program
	var rootCmd = &cobra.Command{Use: "Currency Converter", Short: "A beautiful CLI currency"} // Define the root command

	var cmdConvert = &cobra.Command{ // Define the convert command
		Use:   "convert [amount] [base] [target]",     // Usage pattern
		Short: "Convert currency from base to target", // Short description
		Args:  cobra.ExactArgs(3),                     // Require exactly three arguments
		Run: func(cmd *cobra.Command, args []string) { // Function to run for the command
			amount, err := strconv.ParseFloat(args[0], 64) // Parse the amount as a float64
			if err != nil {                                // Check for errors
				log.Fatalf("Invalid amount: %s", args[0]) // Log fatal error and exit if amount is invalid
			}
			base := strings.ToUpper(args[1])      // Convert base currency to upper case
			target := strings.ToUpper(args[2])    // Convert target currency to upper case
			convertCurrency(amount, base, target) // Call the function to convert the currency
		},
	}

	rootCmd.AddCommand(cmdConvert) // Add the convert command to the root command
	rootCmd.Execute()              // Execute the root command
}

func convertCurrency(amount float64, base, target string) { // Function to convert currency
	url := apiURL + base       // Create the URL by appending the base currency to the API URL
	resp, err := http.Get(url) // Make an HTTP GET request
	if err != nil {            // Check for errors
		log.Fatalf("HTTP request failed: %s", err) // Log fatal error and exit if the request fails
	}
	defer resp.Body.Close() // Ensure the response body is closed

	var result exchangeRateResponse                                    // Declare a variable to store the decoded data
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // Decode the JSON response
		log.Fatalf("JSON decoding failed: %s", err) // Log fatal error and exit if decoding fails
	}

	rate, exists := result.Rates[target] // Get the conversion rate for the target currency
	if !exists {                         // Check if the rate exists
		fmt.Printf("Invalid target currency\n") // Print an error message if it doesn't
		return
	}

	convertedAmount := amount * rate                                           // Calculate the converted amount
	fmt.Printf("💵 %.2f %s = %.2f %s\n", amount, base, convertedAmount, target) // Print the result
}
