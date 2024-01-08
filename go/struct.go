package request

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img,omitempty"`
	Description string `json:"description"`
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
