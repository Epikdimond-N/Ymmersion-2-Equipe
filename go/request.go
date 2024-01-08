package request

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetChar() {
	// Read the JSON file
	file, err := os.ReadFile("nico.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON data into Categories struct
	var data map[string]Categories
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Accessing characters under the "Persos" category
	characters, ok := data["categories"]
	if !ok {
		fmt.Println("No 'categories' found in JSON")
		return
	}

	// Display information about each character
	for _, character := range characters.Persos {
		fmt.Printf("Character ID: %d\n", character.ID)
		if character.Img != "" {
			fmt.Printf("Image URL: %s\n", character.Img)
		}
		fmt.Printf("Name: %s\n", character.Name)
		fmt.Printf("Description: %s\n", character.Description)

		fmt.Println("-------------")
	}
}
