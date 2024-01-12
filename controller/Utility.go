package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	One "onepiece/go"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	logged       bool
	users        = make(map[string]One.User) // Map to store users
	username     string
	password     string
	IsAdmin      bool
	newFieldname string
)

func formatString(input string) string {
	// Remove special characters using regular expression
	regExp := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	processedString := string(regExp.ReplaceAll([]byte(input), []byte("")))

	// Replace spaces with hyphens
	formattedString := strings.ReplaceAll(processedString, " ", "-")

	// Remove common articles like "Le" or "La"
	formattedString = strings.TrimPrefix(formattedString, "le-")
	formattedString = strings.TrimPrefix(formattedString, "la-")

	// Convert to lowercase
	formattedString = strings.ToLower(formattedString)

	return formattedString
}

func retrieveAndProcessImage(w http.ResponseWriter, r *http.Request, fieldName string, fullname string) (string, error) {
	file, handler, err := r.FormFile(fieldName)
	if err != nil {
		http.Error(w, "Error retrieving the "+fieldName, http.StatusInternalServerError)
		return "", err
	}
	defer file.Close()
	ext := filepath.Ext(handler.Filename)

	if fieldName == "PersosImage" {
		newFieldname = "imgpersos"
	}
	if fieldName == "PersosAffiche" {
		newFieldname = "affiches-persos"
	}
	if fieldName == "PersosDrapeau" {
		newFieldname = "drapeaux"
	}
	newFileName := fullname + ext
	path := "/assets/img/" + newFieldname + "/" + newFileName
	path2 := "/static/img" + newFieldname + "/" + newFileName
	// Save the file with the new filename
	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return "", err
	}
	defer dst.Close()

	// Copy the file to the destination directory
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error copying the file", http.StatusInternalServerError)
		return "", err
	}

	return path2, nil
}

func UpdateAdminByUsername(users []One.User, filename string, username string, newAdminValue string) error {
	for i, user := range users {
		if user.Username == username {
			users[i].IsAdmin = newAdminValue

			// Save the updated user data back to the file
			err := SaveUserData(users, filename)
			if err != nil {
				return fmt.Errorf("error saving user data: %v", err)
			}

			return nil
		}
	}
	return fmt.Errorf("user with username %s not found", username)
}

func RetrieveUserData(filename string) ([]One.User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users map[string]One.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	var userList []One.User
	for _, user := range users {
		userList = append(userList, user)
	}

	return userList, nil
}

func searchUser(username string) (One.User, error) {
	// Read the content of the users.json file
	fileContent, err := os.ReadFile("users.json")
	if err != nil {
		return One.User{}, fmt.Errorf("error reading %s: %v", "users.json", err)
	}

	var userMap map[string]One.User
	err = json.Unmarshal(fileContent, &userMap)
	if err != nil {
		return One.User{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	if userData, ok := userMap[username]; ok {
		return userData, nil
	}

	return One.User{}, nil // No error, but user not found
}

func checkAdmin(usernameToCheck string) (string, error) {
	// Read JSON file
	file, err := os.ReadFile("users.json")
	if err != nil {
		return "", fmt.Errorf("can't read file: %v", err)
	}

	// Unmarshal JSON data
	var users map[string]One.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return "", fmt.Errorf("can't unmarshal: %v", err)
	}

	// Check if the username exists
	user, exists := users[usernameToCheck]
	if !exists {
		return "no", nil // Username not found
	}

	// Check if the user has admin privileges
	if user.IsAdmin == "yes" {
		return "yes", nil
	}

	return "no", nil
}

func ChargeImage() {
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
}

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
func SaveUserData(users []One.User, filename string) error {
	userMap := make(map[string]One.User)
	for _, user := range users {
		userMap[user.Username] = user
	}

	data, err := json.MarshalIndent(userMap, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
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

func idExists(data map[string]interface{}, id string) bool {
	categories, ok := data["categories"].(map[string]interface{})
	if !ok {
		return false
	}

	persos, ok := categories["Persos"].([]interface{})
	if !ok {
		return false
	}

	for _, perso := range persos {
		if p, ok := perso.(map[string]interface{}); ok {
			if pID, exists := p["ID"].(string); exists && pID == id {
				return true
			}
		}
	}
	return false
}
func UpdateChar(name string, img string, affiche string, drapeau string, fullname string, prime string, desc string, role string, fruit string, persona string, apparence string, capacite string, histoire string) error {
	// Read JSON data from file
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal the JSON data into a map[string]interface{}
	var parsedData map[string]interface{}
	if err := json.Unmarshal(fileData, &parsedData); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// New character data
	newID := "Persos/" + fullname
	for idExists(parsedData, newID) {
		newID += "/new"
	}

	newPerso := map[string]interface{}{
		"id":      newID,
		"name":    name,
		"img":     img,
		"affiche": affiche,
		"specs": map[string]interface{}{
			"fullName": fullname,
			"prime":    prime,
			"drapeau":  drapeau,
			"aPropos": map[string]string{
				"description": desc,
				"role":        role,
				"demonFruit":  fruit,
				"personalité": persona,
				"apparence":   apparence,
				"capacités":   capacite,
				"histoire":    histoire,
			},
		},
	}

	categories, ok := parsedData["categories"].(map[string]interface{})
	if !ok {
		return errors.New("error accessing categories data")
	}

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
	if err := os.WriteFile("data.json", updatedData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("Successfully added a new perso and updated data.json")
	return nil
}

func UpdateArc(name string, intro string, img string, affiche string, episode string, chapitre string, desc string) error {
	// Read JSON data from file
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal the JSON data into a map[string]interface{}
	var parsedData map[string]interface{}
	if err := json.Unmarshal(fileData, &parsedData); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// New character data
	newID := "Arcs/" + name
	for idExists(parsedData, newID) {
		newID += "/new"
	}

	newArc := map[string]interface{}{
		"id":             newID,
		"name":           name,
		"intro":          intro,
		"img":            img,
		"affiche":        affiche,
		"épisodesAnimé":  episode,
		"chapitresManga": chapitre,
		"description":    desc,
	}

	categories, ok := parsedData["categories"].(map[string]interface{})
	if !ok {
		return errors.New("error accessing categories data")
	}

	arcs, ok := categories["Arcs"].([]interface{})
	if !ok {
		return errors.New("error accessing Persos data")
	}
	arcs = append(arcs, newArc)
	categories["Arcs"] = arcs
	parsedData["categories"] = categories

	// Marshal the modified data back to JSON
	updatedData, err := json.MarshalIndent(parsedData, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Write the updated JSON data back to the file
	if err := os.WriteFile("data.json", updatedData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("Successfully added a new arc and updated data.json")
	return nil
}

func UpdateEvent(name string, affiche string, desc string) error {
	// Read JSON data from file
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal the JSON data into a map[string]interface{}
	var parsedData map[string]interface{}
	if err := json.Unmarshal(fileData, &parsedData); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// New character data
	newID := "Events/" + name
	for idExists(parsedData, newID) {
		newID += "/new"
	}

	newEvent := map[string]interface{}{
		"id":          newID,
		"name":        name,
		"affiche":     affiche,
		"description": desc,
	}

	categories, ok := parsedData["categories"].(map[string]interface{})
	if !ok {
		return errors.New("error accessing categories data")
	}

	events, ok := categories["EventsOnePiece"].([]interface{})
	if !ok {
		return errors.New("error accessing Persos data")
	}
	events = append(events, newEvent)
	categories["EventsOnePiece"] = events
	parsedData["categories"] = categories

	// Marshal the modified data back to JSON
	updatedData, err := json.MarshalIndent(parsedData, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Write the updated JSON data back to the file
	if err := os.WriteFile("data.json", updatedData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("Successfully added a new event and updated data.json")
	return nil
}

func DeletePost(data map[string]interface{}, postType, postID string) error {
	if categories, ok := data["categories"].(map[string]interface{}); ok {
		if posts, ok := categories[postType].([]interface{}); ok {
			for i, post := range posts {
				if postMap, ok := post.(map[string]interface{}); ok {
					if postMap["id"] == postID {
						// Delete the post from the slice
						copy(posts[i:], posts[i+1:])
						data["categories"].(map[string]interface{})[postType] = posts[:len(posts)-1]
						updatedData, err := json.MarshalIndent(data, "", "    ")
						if err != nil {
							return fmt.Errorf("error marshaling JSON: %w", err)
						}

						// Write the updated JSON data back to the file
						if err := os.WriteFile("data.json", updatedData, 0644); err != nil {
							return fmt.Errorf("error writing to file: %w", err)
						}

					}
				}
			}
		}
	}
	return nil
}

// Function to find unique IDs, images, and descriptions based on entity name
func FindInfoByName(search string) []One.SearchResult {
	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Failed to read JSON data:", err)
		return nil
	}

	var categoryData One.CategoryData
	err = json.Unmarshal(jsonData, &categoryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}
	encounteredIDs := make(map[string]bool)
	var searchResults []One.SearchResult // Store the search results directly as One.SearchResult

	// Loop through all categories and search for the name in the description
	for _, characters := range categoryData.Categories {
		for _, character := range characters {
			if strings.Contains(strings.ToLower(character.Specs.Apropos.Description), strings.ToLower(search)) ||
				strings.Contains(strings.ToLower(character.Name), strings.ToLower(search)) ||
				strings.Contains(strings.ToLower(character.Specs.FullName), strings.ToLower(search)) {
				if !encounteredIDs[character.ID] {
					categorie := getCategorieByID(character.ID)
					image := getImageByID(character.ID)
					description := getDescriptionByID(character.ID)
					searchResult := One.SearchResult{
						Affiche:     character.Affiche,
						Categorie:   categorie,
						ID:          character.ID,
						Image:       image,
						Description: description,
					}

					searchResults = append(searchResults, searchResult)
					encounteredIDs[character.ID] = true
				}
			}
		}
	}

	return searchResults
}

func getImageByID(id string) string {
	jsonData, err := os.ReadFile("data.json")
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
	// Read JSON data from file
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	// Unmarshal JSON into the defined structs
	var charactersAndArcs One.Data
	if err := json.Unmarshal(data, &charactersAndArcs); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return ""
	}

	// Access and utilize the unmarshalled data as needed
	//fmt.Println("-----------------------Characters--------------------------")
	for _, character := range charactersAndArcs.Categories.Persos {
		//fmt.Printf("ID: %s, Name: %s\n", character.ID, character.Name)
		//fmt.Println("Specs:", character.Specs)
		//fmt.Println("-------------------------------")
		if id == character.ID {
			return character.Specs.Apropos.Description
		}
	}

	//fmt.Println("----------------------Arcs-----------------------:")
	for _, arc := range charactersAndArcs.Categories.Arcs {
		//fmt.Printf("ID: %s, Name: %s\n", arc.ID, arc.Name)
		//fmt.Println("Description:", arc.Description)
		//fmt.Println("-------------------------------")
		if id == arc.ID {
			return arc.Description
		}
	}
	// Access and utilize the unmarshalled data as needed
	//fmt.Println("--------------One Piece Events ---------------------")
	for _, event := range charactersAndArcs.Categories.EventsOnePiece {
		//fmt.Printf("ID: %s, Name: %s\n", event.ID, event.Name)
		//fmt.Println("Description:", event.Description)
		//fmt.Println("-------------------------------")
		if id == event.ID {
			return event.Description
		}
	}

	return ""
}

func GetRandomItems(items []map[string]interface{}, count int) []map[string]interface{} {
	// Create a new source with a specific seed value
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Shuffle the items
	random.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	// Select 'count' number of items
	if count > len(items) {
		count = len(items)
	}
	return items[:count]
}

func findByID(data One.Data, id string) interface{} {
	for _, character := range data.Categories.Persos {
		if character.ID == id {
			return character
		}
	}
	for _, arc := range data.Categories.Arcs {
		if arc.ID == id {
			return arc
		}
	}
	for _, event := range data.Categories.EventsOnePiece {
		if event.ID == id {
			return event
		}
	}
	return nil // If the ID is not found
}

func getCategorieByID(ID string) string {
	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Failed to read JSON data:", err)
		return ""
	}

	var data One.Data
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	// Check the character category
	for _, character := range data.Categories.Persos {
		if character.ID == ID {
			return "Persos"
		}
	}

	// Check the arc category
	for _, arc := range data.Categories.Arcs {

		if arc.ID == ID {

			return "Arcs"
		}
	}

	// Check the event category
	for _, event := range data.Categories.EventsOnePiece {

		if event.ID == ID {
			return "EventsOnePiece"
		}
	}

	return "aya"
}
