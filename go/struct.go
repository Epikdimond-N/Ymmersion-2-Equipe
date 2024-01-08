package request

type Specs struct {
	FullName    string      `json:"fullName"`
	Age         int         `json:"age"`
	Description Description `json:"description"`
}

type Description struct {
	Role        string `json:"role"`
	Fruit       string `json:"demonFruit"`
	Personalité string `json:"personalité"`
	Apparence   string `json:"apparence"`
	Capacités   string `json:"capacités"`
	Histoire    string `json:"histoire"`
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
