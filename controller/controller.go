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
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := One.CombinedData{
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "newPersos", data)
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
	fullname := r.FormValue("PersosFullName")
	ext := filepath.Ext(handler.Filename)
	newFileName := fullname + ext
	// Create the file in the destination directory
	// Change the file path as per your directory structure
	filePath := filepath.Join("assets", "img", "imgpersos", newFileName)
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
	age, _ := strconv.Atoi(r.FormValue("PersosAge"))
	desc := r.FormValue("PersosDescription")
	role := r.FormValue("PersosRole")
	fruit := r.FormValue("PersosFruit")
	persona := r.FormValue("PersosPersonalite")
	apparence := r.FormValue("PersosApparence")
	capacites := r.FormValue("PersosCapacités")
	histoire := r.FormValue("PersosHistoires")
	ImgPath := "/static/img/imgpersos/" + newFileName
	// Call the function to update character passing the file path as img
	if err := UpdateChar(name, ImgPath, fullname, age, desc, role, fruit, persona, apparence, capacites, histoire); err != nil {
		// Handle error
		http.Error(w, "Error updating character", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Persos?id=Persos/"+fullname, http.StatusFound)
}

func NewArcHandler(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := One.CombinedData{
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "newArc", data)
}

func GestionNewArcHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the multipart form data with a maximum upload size of 10MB
	r.ParseMultipartForm(10 << 20)

	// Retrieve the file from the form
	file, handler, err := r.FormFile("arcImage")
	if err != nil {
		// Handle error
		fmt.Println("Error retrieving the arcImage:", err)
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	name := r.FormValue("arcName")
	ext := filepath.Ext(handler.Filename)
	newFileName := name + ext
	// Create the file in the destination directory
	// Change the file path as per your directory structure
	filePath := filepath.Join("assets", "img", "photoarcs", newFileName)
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

	ImgPath := "/static/img/photoarcs/" + newFileName
	episode := r.FormValue("arcEpisodeAnime")
	chapitre := r.FormValue("arcChapitreManga")
	desc := r.FormValue("arcDescription")

	// Call the function to update arc passing the file path as img
	if err := UpdateArc(name, ImgPath, episode, chapitre, desc); err != nil {
		// Handle error
		http.Error(w, "Error updating character", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Arcs?id=Arcs/"+name, http.StatusFound)
}

func NewEventHandler(w http.ResponseWriter, r *http.Request) {
	//if !logged {
	//	http.Redirect(w, r, "/login", http.StatusSeeOther)
	//	return
	//}

	initTemplate.Temp.ExecuteTemplate(w, "newEvent", nil)
}

func GestionNewEventHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("EventName")
	desc := r.FormValue("EventDescription")

	// Call the function to update event passing the file path as img
	if err := UpdateEvent(name, desc); err != nil {
		// Handle error
		http.Error(w, "Error updating character", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/Events?id=Events/"+name, http.StatusFound)
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
	randomCharacters := GetRandomItems(data.Categories["Persos"], 4)

	// Select 2 random arcs
	randomArcs := GetRandomItems(data.Categories["Arcs"], 4)

	// Select 2 random events
	randomEvents := GetRandomItems(data.Categories["EventsOnePiece"], 4)

	// Create a map to pass selected data to the template
	selectedData := map[string]interface{}{
		"RandomCharacters": randomCharacters,
		"RandomArcs":       randomArcs,
		"RandomEvents":     randomEvents,
	}
	dataS := One.CombinedData{
		Result: selectedData,
		Cat:    username,
		Logged: logged,
	}
	// Pass the selected data to the template for rendering
	initTemplate.Temp.ExecuteTemplate(w, "index", dataS)
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
	dataS := One.CombinedData{
		Result: ToSend,
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "char", dataS)
}

func DisplayArc(w http.ResponseWriter, r *http.Request) {
	// Retrieve the arc ID from the URL query parameter
	arcID := r.URL.Query().Get("id")

	// Check if the ID is empty or not provided
	if arcID == "" {
		http.Error(w, "Arc ID is required", http.StatusBadRequest)
		return
	}
	data := One.GetArcs()
	ToSend, err := One.GetArcByID(data, arcID)
	if err != nil {
		// Handle error (e.g., arc not found)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	dataS := One.CombinedData{
		Result: ToSend,
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "arc", dataS)
}

func DisplayEvent(w http.ResponseWriter, r *http.Request) {
	// Retrieve the event ID from the URL query parameter
	eventID := r.URL.Query().Get("id")

	// Check if the ID is empty or not provided
	if eventID == "" {
		http.Error(w, "Event ID is required", http.StatusBadRequest)
		return
	}
	data := One.GetEvents()
	ToSend, err := One.GetEventByID(data, eventID)
	if err != nil {
		// Handle error (e.g., event not found)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	dataS := One.CombinedData{
		Result: ToSend,
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "event", dataS)
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
	dataS := One.CombinedData{
		Result: persosData,
		Cat:    username,
		Logged: logged,
	}
	// Pass the selected data to the template for rendering
	initTemplate.Temp.ExecuteTemplate(w, "selectchar", dataS)
}
func DisplayArcs(w http.ResponseWriter, r *http.Request) {
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
	persosData, ok := data.Categories["Arcs"]
	if !ok {
		http.Error(w, "Arcs data not found", http.StatusInternalServerError)
		return
	}
	dataS := One.CombinedData{
		Result: persosData,
		Cat:    username,
		Logged: logged,
	}
	// Pass the selected data to the template for rendering
	err = initTemplate.Temp.ExecuteTemplate(w, "selectarc", dataS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DisplayEvents(w http.ResponseWriter, r *http.Request) {
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
	persosData, ok := data.Categories["EventsOnePiece"]
	if !ok {
		http.Error(w, "Events data not found", http.StatusInternalServerError)
		return
	}
	dataS := One.CombinedData{
		Result: persosData,
		Cat:    username,
		Logged: logged,
	}
	// Pass the selected data to the template for rendering
	err = initTemplate.Temp.ExecuteTemplate(w, "selectevent", dataS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DisplayCategories(w http.ResponseWriter, r *http.Request) {
	dataS := One.CombinedData{
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "categories", dataS)
}

func DisplayAdmin(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := One.CombinedData{
		Cat:    username,
		Logged: logged,
		Admin:  IsAdmin,
	}
	initTemplate.Temp.ExecuteTemplate(w, "admin", data)
}

func DisplayAdminAdmin(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !IsAdmin {
		http.Redirect(w, r, "/Home", http.StatusSeeOther)
		return
	}
	data := One.CombinedData{
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "adminadmin", data)
}

func DisplayGestionAdmin(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !IsAdmin {
		http.Redirect(w, r, "/Home", http.StatusSeeOther)
		return
	}
	user := r.FormValue("username")
	userData, err := searchUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := One.CombinedData{
		Result: userData,
		Cat:    username,
		Logged: logged,
	}
	err = initTemplate.Temp.ExecuteTemplate(w, "adminchoix", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DisplaySelectionAdmin(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := One.CombinedData{
		Cat:    username,
		Logged: logged,
		Admin:  IsAdmin,
	}
	initTemplate.Temp.ExecuteTemplate(w, "adminadmin", data)
}
func DisplayGestionSelectionAdmin(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !IsAdmin {
		http.Redirect(w, r, "/Home", http.StatusFound)
	}
	user := r.FormValue("username")
	admin := r.FormValue("admin")
	filename := "users.json"
	users, err := RetrieveUserData(filename)
	if err != nil {
		fmt.Println("Error retrieving user data:", err)
		return
	}

	err = UpdateAdminByUsername(users, filename, user, admin)
	if err != nil {
		fmt.Println("Error updating admin value:", err)
		return
	}

	fmt.Println("Admin value updated and data saved successfully.")

	http.Redirect(w, r, "/Home", http.StatusFound)
}
func DisplayAdminDelete(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !IsAdmin {
		http.Redirect(w, r, "/Home", http.StatusSeeOther)
		return
	}
	data := One.CombinedData{
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "adminDelete", data)
}

func DisplayAdminDeleteConf(w http.ResponseWriter, r *http.Request) {
	if !logged {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !IsAdmin {
		http.Redirect(w, r, "/Home", http.StatusSeeOther)
		return
	}
	ID := r.URL.Query().Get("id")
	filePath := "data.json"

	// Read the JSON file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal the JSON data into the appropriate struct
	var data One.Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}

	result := findByID(data, ID)
	if result != nil {
		fmt.Printf("Data found for ID '%s'\n", ID)
	} else {
		fmt.Printf("No data found for ID '%s'\n", ID)
	}
	Cat := getCategorieByID(ID)
	combinedData := One.CombinedData{
		Result: result,
		Cat:    Cat,
		Cat2:   username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "adminConf", combinedData)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	filePath := "data.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}
	postType := getCategorieByID(ID)
	DeletePost(data, postType, ID)

	http.Redirect(w, r, "/Home", http.StatusFound)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := One.CombinedData{
		Cat:    username,
		Logged: logged,
	}
	initTemplate.Temp.ExecuteTemplate(w, "404", data)
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("name")
	if search == "" {
		http.Error(w, "Please provide a name to search", http.StatusBadRequest)
		return
	}

	searchResults := FindInfoByName(search)
	data := One.CombinedData{
		Result: searchResults,
		Cat:    username,
		Logged: logged,
	}
	// Execute the template with searchResults
	err := initTemplate.Temp.ExecuteTemplate(w, "search", data)
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
	admin := checkAdmin(username)
	user, exists := users[username]
	if !exists || !CheckPasswordHash(password, user.Password) {
		http.Redirect(w, r, "/login?invalid=true", http.StatusSeeOther)
		return
	}
	if admin {
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
