package controller

import (
	"net/http"
	One "onepiece/go"
	initTemplate "onepiece/temp"
)

func DisplayHome(w http.ResponseWriter, r *http.Request) {
	data := One.GetChar()
	initTemplate.Temp.ExecuteTemplate(w, "index", data)
}

func DisplayPersos(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "article", nil)
}

func DisplayArcs(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "article", nil)
}

func DisplayEvents(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "article", nil)
}

func DisplayCategories(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "categories", nil)
}

func DisplaySearch(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "search", nil)
}

func DisplayAdmin(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "admin", nil)
}

func DisplayAddArticle(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "addarticle", nil)
}

func Display404(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "404", nil)
}
