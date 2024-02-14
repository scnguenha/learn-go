package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Base64 encoded string
	base64String := ""

	// Decode the base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}

	// Convert decoded bytes to string
	decodedString := string(decodedBytes)

	fmt.Println("Decoded string:", decodedString)
}
