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

func GetArcs() []Arc {
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

	// Accessing characters under the "Arcs" category
	arcs, ok := data["categories"]
	if !ok {
		fmt.Println("No 'categories' found in JSON")
		return nil
	}
	var arc []Arc
	// Display information about each arcs
	for _, Arcs := range arcs.Arc {
		fmt.Printf("Arcs ID: %d\n", Arcs.ID)
		if Arcs.Img != "" {
			fmt.Printf("Image URL: %s\n", Arcs.Img)
		}
		fmt.Printf("Name: %s\n", Arcs.Name)
		fmt.Printf("Description: %s\n", Arcs.Description)

		fmt.Println("-------------")
		var newArcs Arc
		newArcs.ID = Arcs.ID
		newArcs.Name = Arcs.Name
		if Arcs.Img != "" {
			newArcs.Img = Arcs.Img
		}
		newArcs.Description = Arcs.Description

		arc = append(arc, newArcs)
	}

	return arc
}

func GetEvents() []Event {
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

	// Accessing characters under the "Events" category
	events, ok := data["categories"]
	if !ok {
		fmt.Println("No 'categories' found in JSON")
		return nil
	}
	var event []Event
	// Display information about each event
	for _, Events := range events.Events {
		fmt.Printf("Event ID: %d\n", Events.ID)
		if Events.Img != "" {
			fmt.Printf("Image URL: %s\n", Events.Img)
		}
		fmt.Printf("Name: %s\n", Events.Name)
		fmt.Printf("Description: %s\n", Events.Description)

		fmt.Println("-------------")
		var newEvent Event
		newEvent.ID = Events.ID
		newEvent.Name = Events.Name
		if Events.Img != "" {
			newEvent.Img = Events.Img
		}
		newEvent.Description = Events.Description

		event = append(event, newEvent)
	}

	return event
}
