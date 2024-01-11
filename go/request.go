package request

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetChar() []Character {
	// Read the JSON file
	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Unmarshal JSON data into Data struct
	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	// Accessing characters under the "Persos" category
	characters := data.Categories.Persos

	// Display information about each character
	var char []Character
	for _, character := range characters {
		var newChar Character
		newChar.ID = character.ID
		newChar.Name = character.Name
		if character.Img != "" {
			newChar.Img = character.Img
		}
		newChar.Specs.FullName = character.Specs.FullName
		newChar.Specs.Prime = character.Specs.Prime
		newChar.Specs.Apropos.Description = character.Specs.Apropos.Description
		newChar.Specs.Apropos.Role = character.Specs.Apropos.Role
		newChar.Specs.Apropos.Fruit = character.Specs.Apropos.Fruit
		newChar.Specs.Apropos.Personalité = character.Specs.Apropos.Personalité
		newChar.Specs.Apropos.Apparence = character.Specs.Apropos.Apparence
		newChar.Specs.Apropos.Capacités = character.Specs.Apropos.Capacités
		newChar.Specs.Apropos.Histoire = character.Specs.Apropos.Histoire

		char = append(char, newChar)
	}

	return char
}

func GetArcs() []Arc {
	// Read the JSON file
	file, err := os.ReadFile("data.json")
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
		// fmt.Printf("Arcs ID: %s\n", Arcs.ID)
		// if Arcs.Img != "" {
		// 	fmt.Printf("Image URL: %s\n", Arcs.Img)
		// }
		// fmt.Printf("Name: %s\n", Arcs.Name)
		// fmt.Printf("Description: %s\n", Arcs.Description)
		// fmt.Println("-------------")

		var newArcs Arc
		newArcs.ID = Arcs.ID
		newArcs.Name = Arcs.Name
		if Arcs.Img != "" {
			newArcs.Img = Arcs.Img
		}
		newArcs.Episode = Arcs.Episode
		newArcs.Chapitre = Arcs.Chapitre
		newArcs.Description = Arcs.Description

		arc = append(arc, newArcs)
	}

	return arc
}

func GetEvents() []Event {
	// Read the JSON file
	file, err := os.ReadFile("data.json")
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
		// fmt.Printf("Event ID: %s\n", Events.ID)
		// if Events.Img != "" {
		// 	fmt.Printf("Image URL: %s\n", Events.Img)
		// }
		// fmt.Printf("Name: %s\n", Events.Name)
		// fmt.Printf("Description: %s\n", Events.Description)
		// fmt.Println("-------------")

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

func GetCharacterByID(characters []Character, id string) (Character, error) {
	for _, character := range characters {
		if character.ID == id {
			return character, nil // Return the character if the ID matches
		}
	}
	return Character{}, fmt.Errorf("character with ID %s not found", id) // Return an error if the ID is not found
}

func GetArcByID(arcs []Arc, id string) (Arc, error) {
	for _, arc := range arcs {
		if arc.ID == id {
			return arc, nil // Return the character if the ID matches
		}
	}
	return Arc{}, fmt.Errorf("arc with ID %s not found", id) // Return an error if the ID is not found
}

func GetEventByID(events []Event, id string) (Event, error) {
	for _, event := range events {
		if event.ID == id {
			return event, nil // Return the character if the ID matches
		}
	}
	return Event{}, fmt.Errorf("event with ID %s not found", id) // Return an error if the ID is not found
}
