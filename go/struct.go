package request

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
	Age      int     `json:"age"`
	Apropos  APropos `json:"aPropos"`
}
type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Img   string `json:"img,omitempty"`
	Specs Specs  `json:"specs"`
}

type Arc struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img,omitempty"`
	Episode     string `json:"épisodes"`
	Chapitre    string `json:"chapitre"`
	Description string `json:"description"`
}

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img,omitempty"`
	Description string `json:"description"`
}

type Categories struct {
	Persos []Character `json:"Persos"`
	Arc    []Arc       `json:"Arcs"`
	Events []Event     `json:"Events"`
}

// Login part, plz don't touch >>

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  string `json:"admin"`
}

// Login part, plz don't touch <<
