package request

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetChar() []Character {
	// Read the JSON file
	file, err := os.ReadFile("nico.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Unmarshal JSON data into Categories struct
	var data map[string]Categories
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	// Accessing characters under the "Persos" category
	characters, ok := data["categories"]
	if !ok {
		fmt.Println("No 'categories' found in JSON")
		return nil
	}

	// Display information about each character
	var char []Character
	for _, character := range characters.Persos {
		fmt.Printf("Character ID: %d\n", character.ID)
		if character.Img != "" {
			fmt.Printf("Image URL: %s\n", character.Img)
		}
		fmt.Printf("Name: %s\n", character.Name)
		fmt.Printf("Description: %s\n", character.Description)

		fmt.Println("-------------")
		var newChar Character
		newChar.ID = character.ID
		newChar.Name = character.Name
		if character.Img != "" {
			newChar.Img = character.Img
		}
		newChar.Description = character.Description

		char = append(char, newChar)
	}

	return char
}

func GetArcs() {
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

	// Accessing characters under the "Arcs" category
	arcs, ok := data["categories"]
	if !ok {
		fmt.Println("No 'categories' found in JSON")
		return
	}

	// Display information about each arcs
	for _, Arcs := range arcs.Arc {
		fmt.Printf("Character ID: %d\n", Arcs.ID)
		if Arcs.Img != "" {
			fmt.Printf("Image URL: %s\n", Arcs.Img)
		}
		fmt.Printf("Name: %s\n", Arcs.Name)
		fmt.Printf("Description: %s\n", Arcs.Description)

		fmt.Println("-------------")
	}
}

func GetEvents() {
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

	// Accessing characters under the "Events" category
	events, ok := data["categories"]
	if !ok {
		fmt.Println("No 'categories' found in JSON")
		return
	}

	// Display information about each event
	for _, Events := range events.Events {
		fmt.Printf("Character ID: %d\n", Events.ID)
		if Events.Img != "" {
			fmt.Printf("Image URL: %s\n", Events.Img)
		}
		fmt.Printf("Name: %s\n", Events.Name)
		fmt.Printf("Description: %s\n", Events.Description)

		fmt.Println("-------------")
	}
}
