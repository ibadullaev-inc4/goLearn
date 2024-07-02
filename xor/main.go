package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'chiper' or 'decipher' . Default is 'chiper'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()

	switch *mode {

	case "cipher":
		plaintext := getUserInput("Enter your cipher: ")
		cipheredText, err := cipherer.Cipher(plaintext, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encrypting text: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(cipheredText)
	case "decipher":
		cipheredText := getUserInput("Enter your ciphered data to decipher: ")
		decipheredText, err := cipherer.Decipher(cipheredText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decrypting text: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(decipheredText)

	default:
		fmt.Println("Invalid mode. Use cipher or decipher")
		os.Exit(1)
	}

	if len(*secretKey) == 0 {
		fmt.Fprintln(os.Stderr, "No secret is provided! Exiting now ...")
		os.Exit(1)
	}
}

func getUserInput(msg string) string {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading the entere text! Please try again ...")
			continue
		}

		return strings.TrimRight(result, "\n")
	}

}
