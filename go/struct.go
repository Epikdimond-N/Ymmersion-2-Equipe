package request

type DataHome struct {
	Categories map[string][]map[string]interface{} `json:"categories"`
}

type Data struct {
	Categories struct {
		Persos         []Character `json:"Persos"`
		Arcs           []Arc       `json:"Arcs"`
		EventsOnePiece []Event     `json:"EventsOnePiece"`
	} `json:"categories"`
}
type SearchResult struct {
	Categorie   string
	ID          string
	Image       string
	Description string
}

type CategoryData struct {
	Categories map[string][]Character `json:"categories"`
	Persos     []Character            `json:"Persos"`
	Arcs       []Arc                  `json:"Arcs"`
	Events     []Event                `json:"Events"`
}

type APropos struct {
	Description string `json:"description"`
	Role        string `json:"role"`
	Fruit       string `json:"demonFruit"`
	Personalité string `json:"personalité"`
	Apparence   string `json:"apparence"`
	Capacités   string `json:"capacités"`
	Histoire    string `json:"histoire"`
}

type Specs struct {
	FullName string  `json:"fullName"`
	Prime    string  `json:"prime"`
	Apropos  APropos `json:"aPropos"`
	Drapeau  string  `json:"drapeau"`
}
type Character struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Img     string `json:"img"`
	Affiche string `json:"affiche"`
	Specs   Specs  `json:"specs"`
}

type Arc struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img"`
	Affiche     string `json:"affiche"`
	Episode     string `json:"épisodes"`
	Chapitre    string `json:"chapitre"`
	Description string `json:"description"`
}

type Event struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img"`
	Affiche     string `json:"affiche"`
	Description string `json:"description"`
}

type Categories struct {
	Persos []Character `json:"Persos"`
	Arc    []Arc       `json:"Arcs"`
	Events []Event     `json:"EventsOnePiece"`
}

type ResponseData struct {
	ID          string
	Img         string
	Description string
}

// Login part, plz don't touch >>

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  string `json:"admin"`
}

// Login part, plz don't touch <<
