package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	One "onepiece/go"
	"os"
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
// updating still not done !!!!!
func UpdateChar() error {
	// Read the JSON file
	file, err := os.ReadFile("nico.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Unmarshal JSON data into Categories struct
	var data map[string]One.Categories
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

	fmt.Printf("json:%v\n", characters)

	//----------------------------------------------------------

	var myStruct []One.Character
	var newChar One.Character
	newChar.ID = 99
	newChar.Name = "Nicolas"
	newChar.Img = ""
	newChar.Specs.FullName = "Nicolas D. Moyon"
	newChar.Specs.Age = 27
	newChar.Specs.Apropos.Description = "Fait son code"
	newChar.Specs.Apropos.Role = "Student"
	newChar.Specs.Apropos.Fruit = "le fruit du PC"
	newChar.Specs.Apropos.Personalité = "Fatigué"
	newChar.Specs.Apropos.Apparence = "Beau gosse"
	newChar.Specs.Apropos.Capacités = "Pertinan, a soudainement une bonne idée"
	newChar.Specs.Apropos.Histoire = "Rien a dire, no comment"
	myStruct = append(myStruct, newChar)
	jsonFromSlice, err := json.MarshalIndent(myStruct, "", " ")
	if err != nil {
		fmt.Println("Error Marshaling:", err)
		return nil
	}
	fmt.Println(string(jsonFromSlice))
	return nil
}
