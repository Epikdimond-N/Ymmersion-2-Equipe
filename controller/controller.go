package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	One "onepiece/go"
	initTemplate "onepiece/temp"
	"os"
	"path/filepath"
	"strconv"
)

func NewCharHandler(w http.ResponseWriter, r *http.Request) {
	//if !logged {
	//	http.Redirect(w, r, "/login", http.StatusSeeOther)
	//	return
	//}

	initTemplate.Temp.ExecuteTemplate(w, "newPersos", nil)
}

func GestionNewPersosHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data with a maximum upload size of 10MB
	r.ParseMultipartForm(10 << 20)

	// Retrieve the file from the form
	file, handler, err := r.FormFile("PersosImage")
	if err != nil {
		// Handle error
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create the file in the destination directory
	// Change the file path as per your directory structure
	filePath := filepath.Join("assets", "imgpersos", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		// Handle error
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file to the destination directory
	if _, err = io.Copy(dst, file); err != nil {
		// Handle error
		http.Error(w, "Error copying the file", http.StatusInternalServerError)
		return
	}

	// Once the file is saved, retrieve other form data and call the function to update the character
	name := r.FormValue("PersosName")
	fullname := r.FormValue("PersosFullName")
	age, _ := strconv.Atoi(r.FormValue("PersosAge"))
	desc := r.FormValue("PersosDescription")
	role := r.FormValue("PersosRole")
	fruit := r.FormValue("PersosFruit")
	persona := r.FormValue("PersosPersonalite")
	apparence := r.FormValue("PersosApparence")
	capacites := r.FormValue("PersosCapacités")
	histoire := r.FormValue("PersosHistoires")

	// Call the function to update character passing the file path as img
	if err := UpdateChar(name, filePath, fullname, age, desc, role, fruit, persona, apparence, capacites, histoire); err != nil {
		// Handle error
		http.Error(w, "Error updating character", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Home", http.StatusFound)
}

func NewArcHandler(w http.ResponseWriter, r *http.Request) {
	//if !logged {
	//	http.Redirect(w, r, "/login", http.StatusSeeOther)
	//	return
	//}

	initTemplate.Temp.ExecuteTemplate(w, "newEvent", nil)
}

func GestionNewArcHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/Home", http.StatusFound)
}

func NewEventHandler(w http.ResponseWriter, r *http.Request) {
	//if !logged {
	//	http.Redirect(w, r, "/login", http.StatusSeeOther)
	//	return
	//}

	initTemplate.Temp.ExecuteTemplate(w, "newArc", nil)
}

func GestionNewEventHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/Home", http.StatusFound)
}

func DisplayHome(w http.ResponseWriter, r *http.Request) {
	// Open and read the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode JSON data into a Data struct
	var data One.DataHome
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Select 2 random characters
	randomCharacters := GetRandomItems(data.Categories["Persos"], 2)

	// Select 2 random arcs
	randomArcs := GetRandomItems(data.Categories["Arcs"], 2)

	// Select 2 random events
	randomEvents := GetRandomItems(data.Categories["EventsOnePiece"], 2)

	// Create a map to pass selected data to the template
	selectedData := map[string]interface{}{
		"RandomCharacters": randomCharacters,
		"RandomArcs":       randomArcs,
		"RandomEvents":     randomEvents,
	}

	// Pass the selected data to the template for rendering
	initTemplate.Temp.ExecuteTemplate(w, "index", selectedData)
}

// show only one by ID >>
func DisplayPerso(w http.ResponseWriter, r *http.Request) {
	// Retrieve the character ID from the URL query parameter
	charID := r.URL.Query().Get("id")

	// Check if the ID is empty or not provided
	if charID == "" {
		http.Error(w, "Character ID is required", http.StatusBadRequest)
		return
	}

	data := One.GetChar()
	ToSend, err := One.GetCharacterByID(data, charID)
	if err != nil {
		// Handle error (e.g., character not found)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	initTemplate.Temp.ExecuteTemplate(w, "char", ToSend)
}
func DisplayArc(w http.ResponseWriter, r *http.Request) {
	// Retrieve the character ID from the URL query parameter
	arcID := r.URL.Query().Get("id")

	// Check if the ID is empty or not provided
	if arcID == "" {
		http.Error(w, "Arc ID is required", http.StatusBadRequest)
		return
	}
	data := One.GetArcs()
	ToSend, err := One.GetArcByID(data, arcID)
	if err != nil {
		// Handle error (e.g., character not found)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	initTemplate.Temp.ExecuteTemplate(w, "arc", ToSend)
}
func DisplayEvent(w http.ResponseWriter, r *http.Request) {
	data := One.GetChar()
	ToSend, err := One.GetCharacterByID(data, "1")
	if err != nil {
		// Handle error (e.g., character not found)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	initTemplate.Temp.ExecuteTemplate(w, "char", ToSend)
}

// <<<<

// show multiple by categorie >>
func DisplayPersos(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode JSON data into a DataHome struct
	var data One.DataHome

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the "Persos" data
	persosData, ok := data.Categories["Persos"]
	if !ok {
		http.Error(w, "Persos data not found", http.StatusInternalServerError)
		return
	}

	// Pass the selected data to the template for rendering
	initTemplate.Temp.ExecuteTemplate(w, "selectchar", persosData)
}
func DisplayArcs(w http.ResponseWriter, r *http.Request) {
	// if !logged {
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	return
	// }
	initTemplate.Temp.ExecuteTemplate(w, "article", nil)
}
func DisplayEvents(w http.ResponseWriter, r *http.Request) {
	// if !logged {
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	return
	// }
	initTemplate.Temp.ExecuteTemplate(w, "article", nil)
}

// <<<<

func DisplayCategories(w http.ResponseWriter, r *http.Request) {
	// if !logged {
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	return
	// }
	initTemplate.Temp.ExecuteTemplate(w, "categories", nil)
}

func DisplayAdmin(w http.ResponseWriter, r *http.Request) {
	// if !logged {
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	return
	// }
	// if !admin {
	// 	http.Redirect(w, r, "/Home", http.StatusSeeOther)
	// 	return
	// }
	initTemplate.Temp.ExecuteTemplate(w, "admin", nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "404", nil)
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("name")
	if search == "" {
		http.Error(w, "Please provide a name to search", http.StatusBadRequest)
		return
	}

	searchResults := FindInfoByName(search)

	// Execute the template with searchResults
	err := initTemplate.Temp.ExecuteTemplate(w, "search", searchResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Login part, warning >>
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "Register", nil)
}

func ConfirmRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/Home", http.StatusSeeOther)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	isAdmin := r.FormValue("admin")

	// Check if username already exists
	if _, exists := users[username]; exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	hashedPassword := HashPassword(password)
	users[username] = One.User{Username: username, Password: hashedPassword, IsAdmin: isAdmin}
	// Save users to a file
	if err := saveUsersToFile("users.json"); err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Load users from a file on startup
	if err := LoadUsersFromFile("users.json"); err != nil {
		panic(err)
	}
	// Check if there are query parameters in the URL
	queryParams := r.URL.Query()
	// Get a specific query parameter value by key
	invalidParam := queryParams.Get("invalid")
	var Invalid string
	Invalid = ""
	// Use the obtained query parameter value
	if invalidParam != "" {
		Invalid = "Invalid username or password"
		invalidParam = ""
	}

	initTemplate.Temp.ExecuteTemplate(w, "Login", Invalid)
}

func SuccessLoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/Home", http.StatusSeeOther)
		return
	}

	username = r.FormValue("username")
	password = r.FormValue("password")
	admin := r.FormValue("admin")
	user, exists := users[username]
	if !exists || !CheckPasswordHash(password, user.Password) {
		http.Redirect(w, r, "/login?invalid=true", http.StatusSeeOther)
		return
	}

	if admin == "yes" {
		IsAdmin = true
	}
	logged = true
	// Successfully logged in
	// Handle further operations (e.g., setting session, redirecting, etc.)
	http.Redirect(w, r, "/Home", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ResetUserValue()
	http.Redirect(w, r, "/Home", http.StatusSeeOther)
}

func ChangeLoginHandler(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/Home", http.StatusSeeOther)
		return
	}
	oldpassword := r.FormValue("oldpassword")
	newpassword := r.FormValue("newpassword")
	err := UpdateUserCredentials(username, oldpassword, newpassword)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Password updated successfully.")
	ResetUserValue()
	http.Redirect(w, r, "/Home", http.StatusSeeOther)
}

//Login part , warning <<
