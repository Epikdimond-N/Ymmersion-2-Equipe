package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	One "onepiece/go"
	"os"
	"strings"
)

var (
	logged   bool
	users    = make(map[string]One.User) // Map to store users
	username string
	password string
	IsAdmin  bool
)

func ResetUserValue() {
	logged = false
	username = ""
	password = ""
	IsAdmin = false

}

// Login , warning >>
func LoadUsersFromFile(filename string) error {
	// Check if the file exists
	fileInfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// Create an empty users.json file if it doesn't exist
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}

	// Check if the file is empty
	if fileInfo != nil && fileInfo.Size() == 0 {
		// File exists but is empty, so initialize users as an empty map
		users = make(map[string]One.User)
		return nil
	}

	// Load users from the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Check if the file contains valid JSON data
	if len(data) == 0 {
		// File is empty or doesn't contain valid JSON
		return nil
	}

	users = make(map[string]One.User)
	if err := json.Unmarshal(data, &users); err != nil {
		return err
	}

	return nil
}

func UpdateUserCredentials(name, oldPassword, newPassword string) error {
	// Read the JSON file into memory
	raw, err := os.ReadFile("users.json")
	if err != nil {
		return err
	}

	// Define a struct that matches your JSON structure
	var data map[string]One.User // Map where keys are strings and values are User structs

	// Unmarshal the JSON into the defined struct
	if err := json.Unmarshal(raw, &data); err != nil {
		return err
	}

	// Check if the user exists in the map
	user, exists := data[name]
	if !exists {
		return fmt.Errorf("user not found")
	}

	if !CheckPasswordHash(oldPassword, user.Password) {
		return fmt.Errorf("incorrect password")
	}

	if newPassword != "" {
		// Update the password
		user.Password = HashPassword(newPassword)

		// Update the user in the map
		data[name] = user

		// Marshal the updated data back to JSON
		updatedJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		// Write the updated JSON back to the file
		err = os.WriteFile("users.json", updatedJSON, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// Function to save users to a file for register func
func saveUsersToFile(filename string) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		log.Println("Error writing updated user data:", err)
		return err
	}

	log.Println("User data successfully updated")
	return nil
}

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hasher.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	hashedPassword := HashPassword(password)
	return hashedPassword == hash
}

// Login , warning <<

func UpdateChar(name string, img string, fullname string, age int, desc string, role string, fruit string, persona string, apparence string, capacite string, histoire string) error {
	// Read JSON data from file
	fileData, err := os.ReadFile("nico.json")
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal the JSON data into a map[string]interface{}
	var parsedData map[string]interface{}
	if err := json.Unmarshal(fileData, &parsedData); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// New character data
	newPerso := map[string]interface{}{
		"ID":   99,
		"Name": name,
		"Img":  img,
		"Specs": map[string]interface{}{
			"FullName": fullname,
			"Age":      age,
			"Apropos": map[string]string{
				"Description": desc,
				"Role":        role,
				"Fruit":       fruit,
				"Personalité": persona,
				"Apparence":   apparence,
				"Capacités":   capacite,
				"Histoire":    histoire,
			},
		},
	}

	categories, ok := parsedData["categories"].(map[string]interface{})
	if !ok {
		return errors.New("error accessing categories data")
	}
	// Append the new character to the "Persos" array in parsedData
	persos, ok := categories["Persos"].([]interface{})
	if !ok {
		return errors.New("error accessing Persos data")
	}
	persos = append(persos, newPerso)
	categories["Persos"] = persos
	parsedData["categories"] = categories

	// Marshal the modified data back to JSON
	updatedData, err := json.MarshalIndent(parsedData, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Write the updated JSON data back to the file
	if err := os.WriteFile("nico.json", updatedData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("Successfully added a new perso and updated nico.json")
	return nil

}

// Function to find unique IDs, images, and descriptions based on entity name
func FindInfoByName(search string) []One.SearchResult {
	jsonData, err := os.ReadFile("nico.json")
	if err != nil {
		fmt.Println("Failed to read JSON data")
		return nil
	}

	var categoryData One.CategoryData
	err = json.Unmarshal(jsonData, &categoryData)
	if err != nil {
		fmt.Println("Error parsing JSON")
		return nil
	}

	var searchResults []struct {
		ID       string
		Category string
	}

	encounteredIDs := make(map[string]bool)

	// Loop through all categories and search for the name in the description
	for category, characters := range categoryData.Categories {
		for _, character := range characters {
			if strings.Contains(strings.ToLower(character.Specs.Apropos.Description), strings.ToLower(search)) {
				if !encounteredIDs[character.ID] {
					searchResults = append(searchResults, struct {
						ID       string
						Category string
					}{
						ID:       character.ID,
						Category: category,
					})
					encounteredIDs[character.ID] = true
				}
			}
		}
	}
	fmt.Println(searchResults)

	var searchResults1 []One.SearchResult

	// Loop through search results and populate the slice
	for _, result := range searchResults {
		// Assuming you have functions to get image and description by ID
		image := getImageByID(result.ID)
		description := getDescriptionByID(result.ID)

		searchResult := One.SearchResult{
			ID:          result.ID,
			Image:       image,
			Description: description,
		}

	}

	return searchResults1
}

func getImageByID(id string) string {
	jsonData, err := os.ReadFile("nico.json")
	if err != nil {
		fmt.Println("Failed to read JSON data:", err)
		return ""
	}

	var categoryData One.CategoryData
	err = json.Unmarshal(jsonData, &categoryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	for _, characters := range categoryData.Categories {
		for _, character := range characters {
			if character.ID == id {
				return character.Img
			}
		}
	}

	fmt.Println("Image not found for ID:", id)
	return ""
}

func getDescriptionByID(id string) string {
	jsonData, err := os.ReadFile("nico.json")
	if err != nil {
		fmt.Println("Failed to read JSON data:", err)
		return ""
	}

	var categoryData One.CategoryData
	err = json.Unmarshal(jsonData, &categoryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	for _, characters := range categoryData.Categories {
		for _, character := range characters {
			if character.ID == id {
				return character.Specs.Apropos.Description
			}
		}
	}

	fmt.Println("description not found for ID:", id)
	return ""
}
